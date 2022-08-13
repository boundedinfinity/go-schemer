package enum_format

//go:generate enumer -path=./main.go

type EnumFormat string

const (
	String  EnumFormat = "string"
	Integer EnumFormat = "integer"
)
