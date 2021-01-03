package models

type UseCase struct {
	Name            string        `json:"name" mapstructure:"name"`
	IsNameUpperCase bool          `json:"is_name_upper_case" mapstructure:"is_name_upper_case"`
	Implementors    []Implementor `json:"implementors" mapstructure:"implementors"`
}
