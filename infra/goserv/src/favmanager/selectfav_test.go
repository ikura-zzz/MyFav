package favmanager

import "testing"

func TestSelectFavsByUserid(t *testing.T) {
	favs, err := SelectfavsByUserid(1)
	if err != nil {
		t.Fatalf("failed test%#v", err.Error())
	}

	if favs[0].Category == "" {
		t.Fatalf("failed test やっぱりだめ")
	}
}
