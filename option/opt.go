package option

import (
	"fmt"
	"github.com/devproje/commando"
	"strconv"
	"strings"
)

func extractIndex(name string, args []string) (*int, error) {
	for i, arg := range args {
		if strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "--") {
			var remove = arg
			remove = strings.TrimPrefix(remove, "--")
			remove = strings.TrimPrefix(remove, "-")

			if remove == name {
				return &i, nil
			}
		}
	}

	return nil, fmt.Errorf("--%s option is not defined", name)
}

func ParseString(name string, n *commando.Node) (string, error) {
	index, err := extractIndex(name, n.Commando.Args())
	if err != nil {
		return "", err
	}

	if index == nil {
		return "", nil
	}

	if *index+1 >= len(n.Commando.Args()) {
		return "", fmt.Errorf("--%s option's value is not defined", name)
	}

	return n.Commando.Args()[*index+1], nil
}

func ParseInt(name string, n *commando.Node) (int64, error) {
	index, err := extractIndex(name, n.Commando.Args())
	if err != nil {
		return 0, err
	}

	if index == nil {
		return 0, nil
	}

	if *index+1 >= len(n.Commando.Args()) {
		return 0, fmt.Errorf("--%s option's value is not defined", name)
	}

	arg := n.Commando.Args()[*index+1]
	if strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "--") {
		return 0, fmt.Errorf("--%s option's value is not defined", name)
	}

	return strconv.ParseInt(arg, 0, 64)
}

func ParseFloat(name string, n *commando.Node) (float64, error) {
	index, err := extractIndex(name, n.Commando.Args())
	if err != nil {
		return 0, err
	}

	if index == nil {
		return 0, nil
	}

	if *index+1 >= len(n.Commando.Args()) {
		return 0, fmt.Errorf("--%s option's value is not defined", name)
	}

	arg := n.Commando.Args()[*index+1]
	if strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "--") {
		return 0, fmt.Errorf("--%s option's value is not defined", name)
	}

	return strconv.ParseFloat(arg, 64)
}

func ParseBool(name string, n *commando.Node) (bool, error) {
	index, err := extractIndex(name, n.Commando.Args())
	if err != nil {
		return false, err
	}

	if index == nil {
		return false, nil
	}

	if *index+1 >= len(n.Commando.Args()) {
		return true, nil
	}

	arg := n.Commando.Args()[*index+1]
	if arg == "true" || arg == "false" {
		var parsed bool
		parsed, err = strconv.ParseBool(arg)
		if err != nil {
			return false, err
		}

		return parsed, nil
	}

	return true, nil
}
