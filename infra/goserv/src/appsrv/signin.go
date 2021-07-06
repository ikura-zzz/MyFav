package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ログイン要求
func Signin(engine *gin.Engine) {
	engine.POST("/signin", func(c *gin.Context) {
		/*inputs := GetInputPOST(c)
		if err := selectdb.ChkIdentifier(inputs[0], inputs[1]); err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"errmessage": "ログインID,パスワードが一致しません",
			})
		} else {*/
		c.HTML(http.StatusOK, "listall.html", gin.H{
			"list": template.HTML(genlistall()),
			/*"food": template.HTML(selectdb.GetShokuzaiTD_Food()),
			"seas": template.HTML(selectdb.GetShokuzaiTD_Seas()),
			"powd": template.HTML(selectdb.GetShokuzaiTD_Powd()),
			"swet": template.HTML(selectdb.GetShokuzaiTD_Swet()),
			"froz": template.HTML(selectdb.GetShokuzaiTD_Froz()),
			"frut": template.HTML(selectdb.GetShokuzaiTD_Frut()),*/
		})
		//}
	})
}
