package fav

import (
	"errors"

	"myfav/dbaccessor"
	"myfav/types"
)

// Favmod Favの編集
func Favmod(f types.Fav) error {

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
