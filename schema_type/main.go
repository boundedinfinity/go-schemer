package schema_type

//go:generate enumer -path=./main.go

type SchemaType string

const (
	String   SchemaType = "string"
	Float    SchemaType = "float"
	Double   SchemaType = "double"
	Integer  SchemaType = "integer"
	Byte     SchemaType = "byte"
	Char     SchemaType = "char"
	Short    SchemaType = "short"
	Long     SchemaType = "long"
	Ubyte    SchemaType = "ubyte"
	Ulong    SchemaType = "ulong"
	Ushort   SchemaType = "ushort"
	Object   SchemaType = "object"
	Array    SchemaType = "array"
	Boolean  SchemaType = "boolean"
	Enum     SchemaType = "enum"
	Date     SchemaType = "date"
	Time     SchemaType = "time"
	DateTime SchemaType = "datetime"
	Duration SchemaType = "duration"
)
