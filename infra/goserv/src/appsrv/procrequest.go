package main

import (
	"fmt"
	"html/template"
	"myfav/crtusr"
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
		username := c.Query("username")
		password := c.Query("password")
		fmt.Println(username + password)
		switchlistall(c)
	})
}
func Signout(engine *gin.Engine) {
	engine.GET("/signout", func(c *gin.Context) {
		c.Redirect(303, "/")
	})
}
func Signup(engine *gin.Engine) {
	engine.POST("/signup", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		retype := c.PostForm("retypepassword")
		fmt.Println(username)
		fmt.Println(password)
		fmt.Println(retype)
		if err := crtusr.Useradd(username, password, retype); err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"errmsg": err.Error(),
			})
		}
		c.Redirect(303, "/")
	})
}
func Listall(engine *gin.Engine) {
	engine.GET("/listall", func(c *gin.Context) {
		switchlistall(c)
	})
}
func Listalready(engine *gin.Engine) {
	engine.GET("/listalready", func(c *gin.Context) {
		c.HTML(http.StatusOK, "listalready.html", gin.H{})
	})
}
func Listnow(engine *gin.Engine) {
	engine.GET("/listnow", func(c *gin.Context) {
		c.HTML(http.StatusOK, "listnow.html", gin.H{})
	})
}
func Listwish(engine *gin.Engine) {
	engine.GET("/listwish", func(c *gin.Context) {
		c.HTML(http.StatusOK, "listwish.html", gin.H{})
	})
}
func switchlistall(c *gin.Context) {
	c.HTML(http.StatusOK, "listall.html", gin.H{
		"list": template.HTML(genlistall()),
	})
}
func Crtfav(engine *gin.Engine) {
	engine.POST("/crtfav", func(c *gin.Context) {
		switchlistall(c)
	})
}
