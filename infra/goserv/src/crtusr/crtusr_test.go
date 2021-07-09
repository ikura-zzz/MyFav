package crtusr

import (
	"testing"
)

func TestPasswordValid(t *testing.T) {
	if err := passwordValid("aaa", "aaa"); err != nil {
		t.Fatalf("failed test%#v", err.Error())
	}
}

func TestUseradd(t *testing.T) {
	Useradd("shigeji", "shigeji7614", "shigeji7614")
}
