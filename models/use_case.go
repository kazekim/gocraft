package models

type UseCase struct {
	Name             string        `json:"name" mapstructure:"name"`
	IsNotConvertCase bool          `json:"is_not_convert_case" mapstructure:"is_not_convert_case"`
	Implementors     []Implementor `json:"implementors" mapstructure:"implementors"`
}
