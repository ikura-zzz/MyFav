package main

import (
	"html/template"
	selectdb "myfav/selectDb"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登録済み食材照会画面への遷移
func SelectShokuzai(engine *gin.Engine) {
	engine.GET("/selectShokuzai", func(c *gin.Context) {
		c.HTML(http.StatusOK, "selectShokuzai.html", gin.H{
			"food": template.HTML(selectdb.GetShokuzaiTD_Food()),
			"seas": template.HTML(selectdb.GetShokuzaiTD_Seas()),
			"powd": template.HTML(selectdb.GetShokuzaiTD_Powd()),
			"swet": template.HTML(selectdb.GetShokuzaiTD_Swet()),
			"froz": template.HTML(selectdb.GetShokuzaiTD_Froz()),
			"frut": template.HTML(selectdb.GetShokuzaiTD_Frut()),
		})
	})
}
