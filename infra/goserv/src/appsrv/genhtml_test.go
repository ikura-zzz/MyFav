package main

import (
	"myfav/favmanager"
	"testing"
)

func TestGenhtml(t *testing.T) {
	favs, err := favmanager.SelectfavsByUserid(1)
	if err != nil {
		t.Fatalf("failed test%#v", err.Error())
	}
	html, err := genhtml(favs)
	if err != nil {
		t.Fatalf("failed test%#v", err.Error())
	}
	t.Fatalf("failed test%#v", html)
}
