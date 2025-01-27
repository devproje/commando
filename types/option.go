package types

type OptionData struct {
	Name  string
	Desc  string
	Short []string
	Type  OptionType
}

type OptionType int64

const (
	INTEGER OptionType = iota
	BOOLEAN
	STRING
	FLOAT
)
