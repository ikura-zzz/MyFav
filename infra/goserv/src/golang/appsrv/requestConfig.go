package main

import (
	"html/template"
	"myfav/domain/logger"
	"myfav/domain/session"
	"myfav/domain/user"
	"myfav/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setConfig(engine *gin.Engine) {
	configmenu(engine)
	configUser(engine)
	configName(engine)
	configDelUser(engine)
	configValidChk_name(engine)
	configValidChk_pass(engine)
	configValidChk_deluser(engine)
	changeName(engine)
	changePassword(engine)
	deleteUser(engine)
}

func configmenu(engine *gin.Engine) {
	engine.GET("/config", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			c.HTML(http.StatusOK, "config.html", gin.H{})
		})
	})
}

func configUser(engine *gin.Engine) {
	engine.GET("/configuser", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			c.HTML(http.StatusOK, "reconfPassword.html", gin.H{
				"action": template.HTML("/chgnameform"),
			})
		})
	})
}
func configName(engine *gin.Engine) {
	engine.GET("/configpass", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			c.HTML(http.StatusOK, "reconfPassword.html", gin.H{
				"action": template.HTML("/chgpassform"),
			})
		})
	})
}
func configDelUser(engine *gin.Engine) {
	engine.GET("/configuserdel", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			c.HTML(http.StatusOK, "reconfPassword.html", gin.H{
				"action": template.HTML("/userdeleteform"),
			})
		})
	})
}
func configValidChk_name(engine *gin.Engine) {
	var l logger.Logger = new(logger.Logimp)
	engine.POST("/chgnameform", func(c *gin.Context) {
		username, ok := session.GetSessionValue(c, utils.SessionKeyUser)
		if !ok {
			l.Outlog("Config:can't get username:in configuser")
		}
		inputpass := c.PostForm("password")
		if err := user.Invalidchk(username, inputpass); err != nil {
			l.Outlog("cinfigValidChk_name:" + err.Error())
			transPage(c, func(c *gin.Context) {
				c.HTML(http.StatusOK, "reconfPassword.html", gin.H{
					"action": template.HTML("/chgnameform"),
					"errmsg": template.HTML(err.Error()),
				})
			})
			return
		}
		transPage(c, func(c *gin.Context) {
			c.HTML(http.StatusOK, "changeName.html", gin.H{})
		})
	})
}
func configValidChk_pass(engine *gin.Engine) {
	var l logger.Logger = new(logger.Logimp)
	engine.POST("/chgpassform", func(c *gin.Context) {
		username, ok := session.GetSessionValue(c, utils.SessionKeyUser)
		if !ok {
			l.Outlog("Config:can't get username:in configpass")
		}
		inputpass := c.PostForm("password")
		if err := user.Invalidchk(username, inputpass); err != nil {
			transPage(c, func(c *gin.Context) {
				c.HTML(http.StatusOK, "reconfPassword.html", gin.H{
					"action": template.HTML("/chgpassform"),
					"errmsg": template.HTML(err.Error()),
				})
			})
			return
		}
		transPage(c, func(c *gin.Context) {
			c.HTML(http.StatusOK, "changePassword.html", gin.H{})
		})
	})
}
func configValidChk_deluser(engine *gin.Engine) {
	var l logger.Logger = new(logger.Logimp)
	engine.POST("/userdeleteform", func(c *gin.Context) {
		username, ok := session.GetSessionValue(c, utils.SessionKeyUser)
		if !ok {
			l.Outlog("Config:can't get username:in configuserdel")
		}
		inputpass := c.PostForm("password")
		if err := user.Invalidchk(username, inputpass); err != nil {
			transPage(c, func(c *gin.Context) {
				c.HTML(http.StatusOK, "reconfPassword.html", gin.H{
					"action": template.HTML("/userdeleteform"),
					"errmsg": template.HTML(err.Error()),
				})
			})
			return
		}
		transPage(c, func(c *gin.Context) {
			c.HTML(http.StatusOK, "delUser.html", gin.H{})
		})
	})
}

func changeName(engine *gin.Engine) {
	var l logger.Logger = new(logger.Logimp)
	engine.POST("/chgname", func(c *gin.Context) {
		errtrans := func(errmsg string) {
			transPage(c, func(c *gin.Context) {
				c.HTML(http.StatusOK, "changeName.html", gin.H{
					"errmsg": errmsg,
				})
			})
		}
		currentusername, ok := session.GetSessionValue(c, utils.SessionKeyUser)
		if !ok {
			errtrans(utils.CmnErrmsg)
			return
		}
		newusername := c.PostForm("username")
		if newusername == "" {
			errtrans("ユーザー名を入力してください")
		}

		userid, err := session.GetUserId(c)
		if err != nil {
			l.Outlog("changename:can't get username from session")
			errtrans(utils.CmnErrmsg)
			return
		}
		if err := user.Usernamemod(userid, currentusername, newusername); err != nil {
			l.Outlog(("changeName:usernamemod:") + err.Error())
			errtrans(err.Error())
			return
		}

		session.RemoveSession(c)
		session.CreateNewSession(c, newusername)
		transPage(c, func(c *gin.Context) {
			c.Redirect(303, "/config")
		})
	})
}

func changePassword(engine *gin.Engine) {
	var l logger.Logger = new(logger.Logimp)
	engine.POST("/chgpass", func(c *gin.Context) {
		errtrans := func(errmsg string) {
			transPage(c, func(c *gin.Context) {
				c.HTML(http.StatusOK, "changePassword.html", gin.H{
					"errmsg": errmsg,
				})
			})
		}
		username, ok := session.GetSessionValue(c, utils.SessionKeyUser)
		if !ok {
			errtrans(utils.CmnErrmsg)
		}
		newpass := c.PostForm("password")
		retypepass := c.PostForm("retypepassword")
		if newpass == "" || retypepass == "" {
			errtrans("パスワードが未入力です。")
			return
		}

		userid, err := session.GetUserId(c)
		if err != nil {
			l.Outlog("ChangePassword:GetUserid:" + err.Error())
			errtrans(utils.CmnErrmsg)
		}
		if err := user.Userpassmod(userid, username, newpass, retypepass); err != nil {
			errtrans(err.Error())
		}
		transPage(c, func(c *gin.Context) {
			c.Redirect(303, "/config")
		})
	})
}
func deleteUser(engine *gin.Engine) {
	var l logger.Logger = new(logger.Logimp)
	engine.POST("/deluser", func(c *gin.Context) {
		userid, err := session.GetUserId(c)
		if err != nil {
			l.Outlog("delUser:" + err.Error())
			transPage(c, func(c *gin.Context) {
				c.HTML(http.StatusOK, "delUser.html", gin.H{
					"errmsg": "ユーザー情報が取得できません。",
				})
			})
		}
		if err := user.Userdel(userid); err != nil {
			l.Outlog("DeleteUser:" + err.Error())
			transPage(c, func(c *gin.Context) {
				c.HTML(http.StatusOK, "delUser.html", gin.H{
					"errmsg": "ユーザー削除に失敗しました。",
				})
			})
		}
		session.RemoveSession(c)
		redirectTop(c)
	})
}
