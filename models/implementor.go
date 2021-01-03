package models

type Implementor struct {
	Name       string      `json:"name" mapstructure:"name"`
	Attributes []Attribute `json:"attributes" mapstructure:"attributes"`
	Parameters []Parameter `json:"parameters" mapstructure:"parameters"`
}
