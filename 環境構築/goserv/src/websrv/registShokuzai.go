package main

import (
	"fmt"
	insertdb "myfav/insertDb"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 食材の登録実行
func RegistShokuzai(engine *gin.Engine) {
	engine.GET("/registShokuzai", func(c *gin.Context) {
		inputs := GetInputGET(c)
		output, err := insertdb.InsertDb(inputs)
		if err != nil {
			c.HTML(http.StatusOK, "selectShokuzai.html", gin.H{
				"message": fmt.Sprintf("%s", output),
			})
			return
		}
		c.HTML(http.StatusOK, "selectShokuzai.html", gin.H{})
	})
}
