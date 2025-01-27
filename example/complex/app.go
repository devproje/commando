package main

import (
	"encoding/json"
	"fmt"
	"github.com/devproje/commando"
	"github.com/devproje/commando/option"
	"github.com/devproje/commando/types"
	"os"
)

const FILENAME = "config.json"

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Config struct {
	Users []User `json:"users"`
}

func Load() *Config {
	raw, err := os.ReadFile(FILENAME)
	if err != nil {
		ret := Config{
			Users: make([]User, 0),
		}

		r, _ := json.MarshalIndent(ret, "", "\t")
		_ = os.WriteFile(FILENAME, r, 0644)
	}

	var data Config
	_ = json.Unmarshal(raw, &data)

	return &data
}

func Save(c *Config) error {
	raw, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(FILENAME, raw, 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	cnf := Load()
	command := commando.NewCommando(os.Args[1:])

	command.ComplexRoot("user", "user management command", []commando.Node{
		command.Then("add", "user add", func(n *commando.Node) error {
			var name string
			var err error
			var age int64

			name, err = option.ParseString(*n.MustGetOpt("name"), n)
			if err != nil {
				return err
			}

			age, err = option.ParseInt(*n.MustGetOpt("age"), n)
			if err != nil {
				return err
			}

			ret := User{
				Name: name,
				Age:  int(age),
			}

			cnf.Users = append(cnf.Users, ret)
			fmt.Printf("User added: %s\n", ret.Name)

			return Save(cnf)
		}, types.OptionData{
			Name: "name",
			Desc: "user name",
			Type: types.STRING,
		}, types.OptionData{
			Name: "age",
			Desc: "user age",
			Type: types.INTEGER,
		}),
		command.Then("show", "user print", func(n *commando.Node) error {
			var name string
			var err error

			name, err = option.ParseString(*n.MustGetOpt("name"), n)
			if err != nil {
				return err
			}

			var ret *User
			for _, user := range cnf.Users {
				if user.Name == name {
					ret = &user
				}
			}

			if ret == nil {
				return fmt.Errorf("User not found: %s\n", name)
			}

			fmt.Printf("name: %s, age: %d\n", ret.Name, ret.Age)

			return nil
		}, types.OptionData{
			Name: "name",
			Desc: "user name",
			Type: types.STRING,
		}, types.OptionData{
			Name: "age",
			Desc: "user age",
			Type: types.INTEGER,
		}),
	})

	err := command.Execute()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
