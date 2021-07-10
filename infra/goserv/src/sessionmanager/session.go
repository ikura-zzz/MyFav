package sessionmanager

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const sessionname string = "myfavsession"

func CreateNewSessionStore(engine *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	engine.Use(sessions.Sessions(sessionname, store))
}

func CreateNewSession(c *gin.Context, username string) {
	session := sessions.Default(c)
	session.Set("username", username)
	session.Save()
}
func GetSession(c *gin.Context, key string) (string, bool) {
	session := sessions.Default(c)
	value, ok := session.Get(key).(string)
	return value, ok
}
