package main

import (
	"fmt"
	"github.com/devproje/commando"
	"github.com/devproje/commando/option"
	"github.com/devproje/commando/types"
	"os"
)

func main() {
	command := commando.NewCommando(os.Args[1:])
	command.Register("test", "test command", func(n *commando.Node) error {
		name, err := option.ParseString(n.Opts[0], n)
		if err != nil {
			return err
		}

		age, err := option.ParseInt(n.Opts[1], n)
		if err != nil {
			return err
		}

		adult, err := option.ParseBool(n.Opts[2], n)
		if err != nil {
			return err
		}

		fmt.Println(name, age, adult)

		return nil
	}, types.OptionData{
		Name: "name",
		Desc: "print name",
		Type: types.STRING,
	}, types.OptionData{
		Name: "age",
		Desc: "print age",
		Type: types.INTEGER,
	}, types.OptionData{
		Name: "adult",
		Desc: "print adults",
		Type: types.BOOLEAN,
	})

	err := command.Execute()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
