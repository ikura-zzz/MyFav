package insertdb

import (
	"fmt"
	"testing"
)

func TestInsertDb(t *testing.T) {
	inputs := []string{"たまねぎ", "food", "0.5", "", "2021-03-27", ""}
	if output, err := InsertDb(inputs); err != nil {
		t.Fatalf("failed test%#v", fmt.Sprintf("%s", output))
	}
}
