const cvs = document.getElementById( 'cvs' )
    const cw = cvs.width
    const ch = cvs.height
    const out = document.getElementById( 'out' )
    const oh = out.height
    const ow = out.width

    let ix = 0    // 中心座標
    let iy = 0
    let v = 1.0   // 拡大縮小率
    const img  = new Image()
    img.onload = function( _ev ){   // 画像が読み込まれた
        ix = img.width  / 2
        iy = img.height / 2
        let scl = parseInt( cw / img.width * 100 )
        document.getElementById( 'scal' ).value = scl
        scaling( scl )
    }
    function load_img( _url ){  // 画像の読み込み
        img.src = ( _url ? _url : 'https://1.bp.blogspot.com/-AoQB8vIvlcw/W8BOcXcEQ6I/AAAAAAABPZM/rXNbol90tXcxBZBlXsg__xix03b_F4nqwCLcBGAs/s800/pet_cat_omoi_sleep_man.png' )
    }
    load_img()
    function scaling( _v ) {        // スライダーが変った
        v = parseInt( _v ) * 0.01
        draw_canvas( ix, iy )       // 画像更新
    }

    function draw_canvas( _x, _y ){     // 画像更新
        const ctx = cvs.getContext( '2d' )
        ctx.fillStyle = 'rgb(200, 200, 200)'
        ctx.fillRect( 0, 0, cw, ch )    // 背景を塗る
        ctx.drawImage( img,
            0, 0, img.width, img.height,
            (cw/2)-_x*v, (ch/2)-_y*v, img.width*v, img.height*v,
        )
        ctx.strokeStyle = 'rgba(200, 0, 0, 0.8)'
        ctx.strokeRect( (cw-ow)/2, (ch-oh)/2, ow, oh ) // 赤い枠
    }
    function crop_img(){                // 画像切り取り
        const ctx = out.getContext( '2d' )
        ctx.fillStyle = 'rgb(200, 200, 200)'
        ctx.fillRect( 0, 0, ow, oh )    // 背景を塗る
        ctx.drawImage( img,
            0, 0, img.width, img.height,
            (ow/2)-ix*v, (oh/2)-iy*v, img.width*v, img.height*v,
        )
    }

    let mouse_down = false      // canvas ドラッグ中フラグ
    let sx = 0                  // canvas ドラッグ開始位置
    let sy = 0
    cvs.ontouchstart =
    cvs.onmousedown = function ( _ev ){     // canvas ドラッグ開始位置
        mouse_down = true
        sx = _ev.pageX
        sy = _ev.pageY
        return false // イベントを伝搬しない
    }
    cvs.ontouchend =
    cvs.onmouseout =
    cvs.onmouseup = function ( _ev ){       // canvas ドラッグ終了位置
        if ( mouse_down == false ) return
        mouse_down = false
        draw_canvas( ix += (sx-_ev.pageX)/v, iy += (sy-_ev.pageY)/v )
        return false // イベントを伝搬しない
    }
    cvs.ontouchmove =
    cvs.onmousemove = function ( _ev ){     // canvas ドラッグ中
        if ( mouse_down == false ) return
        draw_canvas( ix + (sx-_ev.pageX)/v, iy + (sy-_ev.pageY)/v )
        return false // イベントを伝搬しない
    }
    cvs.onmousewheel = function ( _ev ){    // canvas ホイールで拡大縮小
        let scl = parseInt( parseInt( document.getElementById( 'scal' ).value ) + _ev.wheelDelta * 0.05 )
        if ( scl < 10  ) scl = 10
        if ( scl > 400 ) scl = 400
        document.getElementById( 'scal' ).value = scl
        scaling( scl )
        return false // イベントを伝搬しない
    }
