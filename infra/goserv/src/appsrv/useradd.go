package main

import (
	"errors"
	"fmt"
	insertdb "myfav/insertDb"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ユーザー追加の実行
func Useradd(engine *gin.Engine) {
	engine.POST("/useradd", func(c *gin.Context) {
		inputs := GetInputPOST(c)
		if len(inputs) != 3 {
			c.HTML(http.StatusOK, "registUser.html", gin.H{
				"message": "全項目が入力されていません。",
			})
			return
		}
		if err := passwordValid(inputs[1], inputs[2]); err != nil {
			c.HTML(http.StatusOK, "registUser.html", gin.H{
				"message": "入力されたパスワードが一致しません",
			})
			return
		}
		if output, err := insertdb.RegistUser(inputs[0], inputs[1]); err != nil {
			c.HTML(http.StatusOK, "registUser.html", gin.H{
				"message": "ユーザー登録時に予期しないエラーが発生しました。" + fmt.Sprintf("%s", output),
			})
			return
		}
		c.HTML(http.StatusOK, "registUser.html", gin.H{
			"message": "登録完了しました。",
		})

	})
}

// パスワード登録時の1つ目と2つ目のパスワード入力が一致しているか判定
func passwordValid(pass1 string, pass2 string) error {
	if pass1 != pass2 {
		return errors.New("registPasswordInvalid")
	}
	return nil
}
