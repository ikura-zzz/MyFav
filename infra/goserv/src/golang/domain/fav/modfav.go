package fav

import (
	"errors"

	"myfav/dbaccessor"
	"myfav/domain/session"
	"myfav/types"

	"github.com/gin-gonic/gin"
)

// Favmod Favの編集
func Favmod(c *gin.Context, f types.Fav) error {
	var err error
	f.Userid, err = session.GetUserId(c)
	if err != nil {
		return errors.New("favmod getuserid:" + err.Error())
	}

	f.Timing = timingconv(f.Timing)
	f.Stars = starsconv(f.Stars)
	f.Openclose = opclconv(f.Openclose)

	if f.Title == "" {
		return errors.New("favadd:title is empty")
	}

	if err := dbaccessor.UpdateFav(f); err != nil {
		return errors.New("favadd updateFav:" + err.Error())
	}
	return nil
}
