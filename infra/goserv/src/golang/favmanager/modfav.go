package favmanager

import (
	"errors"

	"myfav/dbaccessor"
	"myfav/domain/fav"
	"myfav/sessionmanager"
	"myfav/types"

	"github.com/gin-gonic/gin"
)

// Favmod Favの編集
func Favmod(c *gin.Context, f types.Fav) error {
	var err error
	f.Userid, err = sessionmanager.GetUserId(c)
	if err != nil {
		return errors.New("favmod getuserid:" + err.Error())
	}

	f.Timing = fav.Timingconv(f.Timing)
	f.Stars = fav.Starsconv(f.Stars)
	f.Openclose = fav.Opclconv(f.Openclose)

	if f.Title == "" {
		return errors.New("favadd:title is empty")
	}

	if err := dbaccessor.UpdateFav(f); err != nil {
		return errors.New("favadd updateFav:" + err.Error())
	}
	return nil
}
