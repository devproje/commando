package commando_test

import (
	"fmt"
	"testing"

	"github.com/devproje/commando"
)

func TestSimple(t *testing.T) {
	command := commando.NewCommando([]string{"test"})
	command.Root("test", "test command", func(n *commando.Node) error {
		fmt.Println("Hello, World!")
		return nil
	})

	err := command.Execute()
	if err != nil {
		t.Errorf("simple command task failed: %v", err)
	}
}

func TestComplex(t *testing.T) {
	command := commando.NewCommando([]string{"test", "apply"})
	command.ComplexRoot("test", "test command", []commando.Node{
		command.Then("apply", "test command", func(n *commando.Node) error {
			fmt.Println("apply task successful!")
			return nil
		}),
	})

	err := command.Execute()
	if err != nil {
		t.Errorf("complex command task failed: %v", err)
	}
}
