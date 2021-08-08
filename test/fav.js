//ここからpopup関連
const cvs = document.getElementById('cvs')
let cw = cvs.width;
let ch = cvs.height;
let out = document.getElementById('canvas')
const oh = 300
const ow = 300

let ix = 0 // 中心座標
let iy = 0
let v = 1.0 // 拡大縮小率
const image = new Image()

function load_img() { // 画像の読み込み
    console.log("a");
    ix = image.width / 2;
    iy = image.height / 2;
    let scl = parseInt(cw / image.width * 100);
    document.getElementById('scal').value = scl;
    scaling(scl);
}

function scaling(_v) { // スライダーが変った
    v = parseInt(_v) * 0.01
    draw_canvas(ix, iy) // 画像更新
}

function draw_canvas(_x, _y) { // 画像更新
    console.log("x:" + _x + "y:" + _y);
    const ctx = cvs.getContext('2d');
    ctx.clearRect(0, 0, ctx.width, ctx.height);
    ctx.fillStyle = 'rgb(255, 255, 255)'
    ctx.fillRect(0, 0, cw, ch) // 背景を塗る
    ctx.drawImage(image,
        0, 0, image.width, image.height,
        (cw / 2) - _x * v, (ch / 2) - _y * v, image.width * v, image.height * v,
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
    ctx.drawImage(image,
        0, 0, image.width, image.height,
        (ow / 2) - ix * v, (oh / 2) - iy * v, image.width * v, image.height * v,
    )
    base64 = out.get(0).toDataURL('image/jpeg');
    document.getElementById("iconimg").src = "";
    document.getElementById("icon").value = base64;
}

let mouse_down = false // canvas ドラッグ中フラグ
let sx = 0 // canvas ドラッグ開始位置
let sy = 0
cvs.ontouchstart = function(_ev) { // canvas ドラッグ開始位置
    console.log("b")
    mouse_down = true
    touchObj = _ev.changedTouches[0]
    sx = parseInt(touchObj.clientX)
    sy = parseInt(touchObj.clientY)
    return false // イベントを伝搬しない
}
cvs.onmousedown = function(_ev) { // canvas ドラッグ開始位置
    console.log("b")
    mouse_down = true
    sx = parseInt(_ev.clientX)
    sy = parseInt(_ev.clientY)
    return false // イベントを伝搬しない
}
cvs.ontouchend = function(_ev) { // canvas ドラッグ終了位置
    console.log("c")
    if (mouse_down == false) return
    mouse_down = false
    touchObj = _ev.changedTouches[0]
    draw_canvas(ix += (sx - touchObj.pageX) / v, iy += (sy - touchObj.pageY) / v)
    return false // イベントを伝搬しない
}
cvs.onmouseout =
    cvs.onmouseup = function(_ev) { // canvas ドラッグ終了位置
        console.log("c")
        if (mouse_down == false) return
        mouse_down = false
        draw_canvas(ix += (sx - _ev.pageX) / v, iy += (sy - _ev.pageY) / v)
        return false // イベントを伝搬しない
    }
cvs.ontouchmove = function(_ev) { // canvas ドラッグ中
    if (mouse_down == false) return
    touchObj = _ev.changedTouches[0]
    draw_canvas(ix + (sx - touchObj.pageX) / v, iy + (sy - touchObj.pageY) / v)
    return false // イベントを伝搬しない
}
cvs.onmousemove = function(_ev) { // canvas ドラッグ中
    if (mouse_down == false) return
    draw_canvas(ix + (sx - _ev.pageX) / v, iy + (sy - _ev.pageY) / v)
    return false // イベントを伝搬しない
}
cvs.onwheel = function(_ev) { // canvas ホイールで拡大縮小
    let scl = parseInt(parseInt(document.getElementById('scal').value) + _ev.deltaY * 0.05)
    if (scl < 10) scl = 10
    if (scl > 400) scl = 400
    document.getElementById('scal').value = scl
    scaling(scl)
    return false // イベントを伝搬しない
}

function popupClose() {
    var popup = document.getElementById('js-popup');
    if (!popup) return;

    var blackBg = document.getElementById('js-black-bg');
    var cropBtn = document.getElementById('crop');
    var closeBtn = document.getElementById('js-close-btn');

    closePopUp(blackBg);
    closePopUp(cropBtn);
    closePopUp(closeBtn);

    function closePopUp(elem) {
        if (!elem) return;
        elem.addEventListener('click', function() {
            popup.classList.toggle('is-show');
            $('input[type=file]').val('');
        });
    }
}
popupClose();

function popupOpen() {
    var popup = document.getElementById('js-popup');
    if (!popup) return;

    popup.classList.toggle('is-show');
}

function drawimage_popup(image) {
    const THUMBNAIL_MAX_HEIGHT = 500;
    const THUMBNAIL_MAX_WIDTH = 500;
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

    //ctx.clearRect(0, 0, width, height);
    load_img();
}

//ここからファイル選択関連
var head = document.getElementsByTagName('head');
var script = document.createElement('script');
script.setAttribute('src', 'https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js');
script.setAttribute('type', 'text/javascript');
script.addEventListener('load', function() {
    $(function() {
        $('input[type=file]').change(function() {
            var file = null;

            // ファイルを取得する
            file = $(this).prop('files')[0];

            // 選択されたファイルがjpeg/png画像でなければ何もせず終了
            if (file.type != 'image/jpeg' && file.type != 'image/png') {
                file = null;
                return;
            }
            // 画像をリサイズする
            var reader = new FileReader();
            reader.onload = function(e) {
                image.onload = function() {
                    drawimage_popup(image);
                    popupOpen();
                }
                image.src = e.target.result;
            }
            reader.readAsDataURL(file);
        });
        // submitが押されたらfile属性は空にしておく。
        $('#create').click(function() {
            $('input[type=file]').val('');
        });
    });
})
document.head.appendChild(script)