package selectdb

import (
	"errors"
	"fmt"
	"testing"
)

func TestmainSuccess(t *testing.T) {
	err := ChkIdentifier("shigeji", "test")
	if err != nil {
		t.Fatalf("faled test%#v", errors.New("error"))
	}
	fmt.Println("ok")
}

func TestExampleFaild(t *testing.T) {
	err := ChkIdentifier("shigeji", "hoge")
	if err != nil {
		t.Fatalf("faled test%#v", errors.New("error"))
	}
	fmt.Println("ok")
}
