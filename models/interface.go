package models

type Interface struct {
	Name         string        `json:"name" mapstructure:"name"`
	Implementors []Implementor `json:"implementors" mapstructure:"implementors"`
}
