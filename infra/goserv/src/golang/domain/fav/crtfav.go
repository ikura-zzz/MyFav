package fav

import (
	"errors"
	"myfav/dbaccessor"
	"myfav/types"
)

// Favadd Favの新規登録
func Favadd(f types.Fav) error {
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
