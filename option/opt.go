package option

import (
	"fmt"
	"github.com/devproje/commando"
	"github.com/devproje/commando/types"
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

func ParseString(opt types.OptionData, n *commando.Node) (string, error) {
	if opt.Type != types.STRING {
		return "", fmt.Errorf("--%s is not a string", opt.Name)
	}

	index, err := extractIndex(opt.Name, n.Commando.Args())
	if err != nil {
		return "", err
	}

	if index == nil {
		return "", nil
	}

	if *index+1 >= len(n.Commando.Args()) {
		return "", fmt.Errorf("--%s option's value is not defined", opt.Name)
	}

	return n.Commando.Args()[*index+1], nil
}

func ParseInt(opt types.OptionData, n *commando.Node) (int64, error) {
	if opt.Type != types.INTEGER {
		return 0, fmt.Errorf("--%s is not a integer", opt.Name)
	}

	index, err := extractIndex(opt.Name, n.Commando.Args())
	if err != nil {
		return 0, err
	}

	if index == nil {
		return 0, nil
	}

	if *index+1 >= len(n.Commando.Args()) {
		return 0, fmt.Errorf("--%s option's value is not defined", opt.Name)
	}

	arg := n.Commando.Args()[*index+1]
	if strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "--") {
		return 0, fmt.Errorf("--%s option's value is not defined", opt.Name)
	}

	return strconv.ParseInt(arg, 0, 64)
}

func ParseFloat(opt types.OptionData, n *commando.Node) (float64, error) {
	if opt.Type != types.FLOAT {
		return 0, fmt.Errorf("--%s is not a string", opt.Name)
	}

	index, err := extractIndex(opt.Name, n.Commando.Args())
	if err != nil {
		return 0, err
	}

	if index == nil {
		return 0, nil
	}

	if *index+1 >= len(n.Commando.Args()) {
		return 0, fmt.Errorf("--%s option's value is not defined", opt.Name)
	}

	arg := n.Commando.Args()[*index+1]
	if strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "--") {
		return 0, fmt.Errorf("--%s option's value is not defined", opt.Name)
	}

	return strconv.ParseFloat(arg, 64)
}

func ParseBool(opt types.OptionData, n *commando.Node) (bool, error) {
	if opt.Type != types.BOOLEAN {
		return false, fmt.Errorf("--%s is not a bool", opt.Name)
	}

	index, err := extractIndex(opt.Name, n.Commando.Args())
	if err != nil {
		return false, nil
	}

	if *index+1 >= len(n.Commando.Args()) {
		return true, nil
	}

	arg := n.Commando.Args()[*index+1]
	if arg == "true" || arg == "false" || arg == "1" || arg == "0" {
		var parsed bool
		parsed, err = strconv.ParseBool(arg)
		if err != nil {
			return false, err
		}

		return parsed, nil
	}

	return true, nil
}
