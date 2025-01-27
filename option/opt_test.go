package option_test

import (
	"fmt"
	"testing"

	"github.com/devproje/commando"
	"github.com/devproje/commando/option"
	"github.com/devproje/commando/types"
)

func TestOption(t *testing.T) {
	args := []string{"test", "--alg", "sha256", "-d"}
	command := commando.NewCommando(args)

	command.Root("test", "test command", func(n *commando.Node) error {
		var alg string
		var debug bool
		var err error

		alg, err = option.ParseString(*n.MustGetOpt("alg"), n)
		if err != nil {
			return err
		}

		debug, err = option.ParseBool(*n.MustGetOpt("debug"), n)
		if err != nil {
			return err
		}

		fmt.Printf("alg: %s, debug: %v\n", alg, debug)
		return nil
	}, types.OptionData{
		Name: "alg",
		Desc: "test option",
		Type: types.STRING,
	}, types.OptionData{
		Name:  "debug",
		Desc:  "test option",
		Type:  types.BOOLEAN,
		Short: []string{"-d"},
	})

	if err := command.Execute(); err != nil {
		t.Errorf("option task failed: %v", err)
	}
}
