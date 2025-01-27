//go:build exclude

package main

import (
	"fmt"
	"os"

	"github.com/devproje/commando"
	"github.com/devproje/commando/option"
	"github.com/devproje/commando/types"
)

func main() {
	// exclude first filename
	command := commando.NewCommando(os.Args[1:])
	command.Root("user", "print user info", func(n *commando.Node) error {
		name, err := option.ParseString(*n.MustGetOpt("name"), n)
		if err != nil {
			return err
		}

		age, err := option.ParseInt(*n.MustGetOpt("age"), n)
		if err != nil {
			return err
		}

		fmt.Printf("name: %s, age: %d\n", name, age)
		return nil
	}, types.OptionData{
		Name: "name",
		Desc: "print user name",
		Type: types.STRING,
	}, types.OptionData{
		Name: "age",
		Desc: "print user age",
		Type: types.INTEGER,
	})

	err := command.Execute()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
