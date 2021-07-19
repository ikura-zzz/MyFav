var head = document.getElementsByTagName('head')
var script = document.createElement('script')
script.setAttribute('src', 'https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js')
script.setAttribute('type', 'text/javascript')
script.addEventListener('load', function() {
    $(function() {
        const THUMBNAIL_MAX_WIDTH = 300;
        const THUMBNAIL_MAX_HEIGHT = 300;
        $('input[type=file]').change(function() {
            var file = null;
            var blob = null;

            // ファイルを取得する
            file = $(this).prop('files')[0];

            // 選択されたファイルがjpeg/png画像でなければ何もせず終了
            if (file.type != 'image/jpeg' && file.type != 'image/png') {
                file = null;
                blob = null;
                return;
            }

            // 画像をリサイズする
            var image = new Image();
            var reader = new FileReader();
            reader.onload = function(e) {
                image.onload = function() {

                    // 縮小後のサイズを計算する
                    var width, height;
                    if (image.width > image.height) {
                        var ratio = image.height / image.width;
                        width = THUMBNAIL_MAX_WIDTH;
                        height = THUMBNAIL_MAX_WIDTH * ratio;
                    } else {
                        var ratio = image.width / image.height;
                        width = THUMBNAIL_MAX_HEIGHT * ratio;
                        height = THUMBNAIL_MAX_HEIGHT;
                    }

                    // 縮小画像を描画するcanvasのサイズを上で算出した値に変更する
                    var canvas = $('#canvas')
                        .attr('width', width)
                        .attr('height', height);

                    var ctx = canvas[0].getContext('2d');

                    // canvasに既に描画されている画像があればそれを消す
                    ctx.clearRect(0, 0, width, height);

                    document.getElementById("iconimg").style.display = "none";
                    // canvasに縮小画像を描画する
                    ctx.drawImage(image,
                        0, 0, image.width, image.height,
                        0, 0, width, height
                    );

                    // canvasから画像をbase64として取得してhidden属性に添付する
                    var base64 = canvas.get(0).toDataURL('image/jpeg');
                    document.getElementById("icon").value = base64
                }
                image.src = e.target.result;
            }
            reader.readAsDataURL(file);
        });
        // submitが押されたらfile属性は空にしておく。
        $('#submit').click(function() {
            $('input[type=file]').val('');
        });
    })
})
document.head.appendChild(script)
