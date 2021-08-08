//ここからpopup関連
const cvs = document.getElementById('cvs')
let cw = cvs.width;
let ch = cvs.height;
let out = document.getElementById('canvas')
const oh = 350
const ow = 350

let ix = 0 // 中心座標
let iy = 0
let v = 1.0 // 拡大縮小率
const img = new Image()

function load_img(imgsrc) { // 画像の読み込み
    img.src = imgsrc;
    ix = img.width / 2;
    iy = img.height / 2;
    let scl = parseInt(cw / img.width * 100);
    document.getElementById('scal').value = scl;
    scaling(scl);
    console.log("a")
}

function scaling(_v) { // スライダーが変った
    v = parseInt(_v) * 0.01
    draw_canvas(ix, iy) // 画像更新
    console.log("b")
}

function draw_canvas(_x, _y) { // 画像更新
    const ctx = cvs.getContext('2d')
        //ctx.clearRect(0, 0, ctx.width, ctx.height);
    console.log("c");
    ctx.fillStyle = 'rgb(255, 255, 255)'
    ctx.fillRect(0, 0, cw, ch) // 背景を塗る
    ctx.drawImage(img,
        0, 0, img.width, img.height,
        (cw / 2) - _x * v, (ch / 2) - _y * v, img.width * v, img.height * v,
    )
    ctx.strokeStyle = 'rgba(200, 0, 0, 0.8)'
    ctx.strokeRect((cw - ow) / 2, (ch - oh) / 2, ow, oh) // 赤い枠
}

function crop_img() { // 画像切り取り
    out = $('#canvas')
        .attr('width', ow)
        .attr('height', oh);
    const ctx = out[0].getContext('2d')
    ctx.fillStyle = 'rgb(255, 255, 255)'
    ctx.fillRect(0, 0, cw, ch) // 背景を塗る
    ctx.drawImage(img,
        0, 0, img.width, img.height,
        (ow / 2) - ix * v, (oh / 2) - iy * v, img.width * v, img.height * v,
    )
    base64 = out.get(0).toDataURL('image/jpeg');
    document.getElementById("iconimg").src = "";
    document.getElementById("icon").value = base64;
}

let mouse_down = false // canvas ドラッグ中フラグ
let sx = 0 // canvas ドラッグ開始位置
let sy = 0
cvs.ontouchstart =
    cvs.onmousedown = function(_ev) { // canvas ドラッグ開始位置
        mouse_down = true
        sx = _ev.pageX
        sy = _ev.pageY
        return false // イベントを伝搬しない
    }
cvs.ontouchend =
    cvs.onmouseout =
    cvs.onmouseup = function(_ev) { // canvas ドラッグ終了位置
        if (mouse_down == false) return
        mouse_down = false
        draw_canvas(ix += (sx - _ev.pageX) / v, iy += (sy - _ev.pageY) / v)
        return false // イベントを伝搬しない
    }
cvs.ontouchmove =
    cvs.onmousemove = function(_ev) { // canvas ドラッグ中
        if (mouse_down == false) return
        draw_canvas(ix + (sx - _ev.pageX) / v, iy + (sy - _ev.pageY) / v)
        return false // イベントを伝搬しない
    }
cvs.onmousewheel = function(_ev) { // canvas ホイールで拡大縮小
    let scl = parseInt(parseInt(document.getElementById('scal').value) + _ev.wheelDelta * 0.05)
    if (scl < 10) scl = 10
    if (scl > 400) scl = 400
    document.getElementById('scal').value = scl
    scaling(scl)
    return false // イベントを伝搬しない
}

function popupImage() {
    var popup = document.getElementById('js-popup');
    if (!popup) return;

    var blackBg = document.getElementById('js-black-bg');
    var cropBtn = document.getElementById('crop');
    var closeBtn = document.getElementById('js-close-btn');
    var showBtn = document.getElementById('js-show-popup');

    closePopUp(blackBg);
    closePopUp(cropBtn);
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

//ここからファイル選択関連
var head = document.getElementsByTagName('head');
var script = document.createElement('script');
script.setAttribute('src', 'https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js');
script.setAttribute('type', 'text/javascript');
script.addEventListener('load', function() {
    $(function() {
        const THUMBNAIL_MAX_WIDTH = 500;
        const THUMBNAIL_MAX_HEIGHT = 500;
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
            var base64;
            reader.onload = function(e) {
                image.onload = function() {

                    // 縮小後のサイズを計算する
                    var width, height, drawwidth, drawheight;
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
                    var canvas = $('#cvs')
                        .attr('width', THUMBNAIL_MAX_WIDTH)
                        .attr('height', THUMBNAIL_MAX_HEIGHT);

                    var ctx = canvas[0].getContext('2d');

                    // canvasに既に描画されている画像があればそれを消す
                    ctx.clearRect(0, 0, width, height);

                    // canvasに縮小画像を描画する
                    ctx.drawImage(image,
                        0, 0, image.width, image.height,
                        0, 0, width, height
                    );

                    // canvasから画像をbase64として取得してhidden属性に添付する
                    base64 = canvas.get(0).toDataURL('image/jpeg');
                    //ctx.clearRect(0, 0, width, height);
                    load_img(base64);
                }
                image.src = e.target.result;
            }
            reader.readAsDataURL(file);
            popupOpen();
        });
        // submitが押されたらfile属性は空にしておく。
        $('#create').click(function() {
            $('input[type=file]').val('');
        });
    });
})
document.head.appendChild(script)