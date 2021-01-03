package models

const (
	ConfigVariableTypeString  ConfigVariableType = "string"
	ConfigVariableTypeInt     ConfigVariableType = "int"
	ConfigVariableTypeFloat64 ConfigVariableType = "float64"
)

type ConfigVariableType string

func (t ConfigVariableType) String() string {
	return string(t)
}
