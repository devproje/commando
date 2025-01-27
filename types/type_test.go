package types_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/devproje/commando/types"
)

func TestOptionType(t *testing.T) {
	str := fmt.Sprint(types.INTEGER, types.BOOLEAN, types.STRING, types.FLOAT)
	if strings.Compare(str, "0 1 2 3") != 0 {
		t.Errorf("enum type is not matches!")
	}
}
