package fav

import (
	"errors"
	"myfav/dbaccessor"
	"myfav/domain/session"
	"myfav/types"

	"github.com/gin-gonic/gin"
)

// Favadd Favの新規登録
func Favadd(c *gin.Context, f types.Fav) error {
	var err error
	f.Userid, err = session.GetUserId(c)
	if err != nil {
		return errors.New("favadd getuserid:" + err.Error())
	}

	f.Timing = timingconv(f.Timing)
	f.Stars = starsconv(f.Stars)
	f.Openclose = opclconv(f.Openclose)

	if f.Title == "" {
		return errors.New("favadd:title is empty")
	}

	if err := dbaccessor.ResistFav(f); err != nil {
		return errors.New("favadd resistFav:" + err.Error())
	}
	return nil
}
