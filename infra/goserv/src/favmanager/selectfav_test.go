package favmanager

import (
	"myfav/utils"
	"testing"
)

func TestSelectFavsByUserid(t *testing.T) {
	favs, err := SelectfavsByUserid(1, utils.SelectFavsByUserid)
	if err != nil {
		t.Fatalf("failed test%#v", err.Error())
	}

	if favs[0].Category == "" {
		t.Fatalf("failed test やっぱりだめ")
	}
}
