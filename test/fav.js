var head = document.getElementsByTagName('head');
var script = document.createElement('script');
script.setAttribute('src', 'https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js');
script.setAttribute('type', 'text/javascript');
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
            popupOpen();
        });
        // submitが押されたらfile属性は空にしておく。
        $('#create').click(function() {
            $('input[type=file]').val('');
        });
    });
})

function imageresize() {
    // 画像をリサイズする
    var image = new Image();
    var reader = new FileReader();
    reader.onload = function(e) {
        image.onload = function() {

            // 縮小後のサイズを計算する
            var width, height, drawwidth, drawheight;
            if (image.width > image.height) {
                width = THUMBNAIL_MAX_WIDTH;
                height = THUMBNAIL_MAX_WIDTH;
                drawwidth = image.height;
                drawheight = image.height;
            } else {
                width = THUMBNAIL_MAX_HEIGHT;
                height = THUMBNAIL_MAX_HEIGHT;
                drawwidth = image.width;
                drawheight = image.width;
            }

            // 縮小画像を描画するcanvasのサイズを上で算出した値に変更する
            var canvas = $('#canvas')
                .attr('width', width)
                .attr('height', height);

            var ctx = canvas[0].getContext('2d');

            // canvasに既に描画されている画像があればそれを消す
            ctx.clearRect(0, 0, width, height);

            ctx.beginPath();
            ctx.arc(150, 150, 150, 0, Math.PI * 2, false);
            ctx.clip();

            document.getElementById("iconimg").style.display = "none";
            // canvasに縮小画像を描画する
            ctx.drawImage(image, 0, 0, drawwidth, drawheight, 0, 0, width, height);

            // canvasから画像をbase64として取得してhidden属性に添付する
            var base64 = canvas.get(0).toDataURL('image/jpeg');
            document.getElementById("iconimg").src = "";
            document.getElementById("icon").value = base64;
        }
        image.src = e.target.result;
    }
    reader.readAsDataURL(file);
}

function popupImage() {
    var popup = document.getElementById('js-popup');
    if (!popup) return;

    var blackBg = document.getElementById('js-black-bg');
    var closeBtn = document.getElementById('js-close-btn');
    var showBtn = document.getElementById('js-show-popup');

    closePopUp(blackBg);
    closePopUp(closeBtn);
    closePopUp(showBtn);

    function closePopUp(elem) {
        if (!elem) return;
        elem.addEventListener('click', function() {
            popup.classList.toggle('is-show');
        });
    }
}
popupImage();

function popupOpen() {
    var popup = document.getElementById('js-popup');
    if (!popup) return;

    popup.classList.toggle('is-show');
}
document.head.appendChild(script)