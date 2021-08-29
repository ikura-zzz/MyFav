package favmanager

import (
	"errors"
	"myfav/dbaccessor"
	"myfav/domain/fav"
	"myfav/sessionmanager"
	"myfav/types"

	"github.com/gin-gonic/gin"
)

// Favadd Favの新規登録
func Favadd(c *gin.Context, f types.Fav) error {
	var err error
	f.Userid, err = sessionmanager.GetUserId(c)
	if err != nil {
		return errors.New("favadd getuserid:" + err.Error())
	}

	f.Timing = fav.Timingconv(f.Timing)
	f.Stars = fav.Starsconv(f.Stars)
	f.Openclose = fav.Opclconv(f.Openclose)

	if f.Title == "" {
		return errors.New("favadd:title is empty")
	}

	if err := dbaccessor.ResistFav(f); err != nil {
		return errors.New("favadd resistFav:" + err.Error())
	}
	return nil
}
