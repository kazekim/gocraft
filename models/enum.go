package models

type Enum struct {
	Name string `json:"name" mapstructure:"name"`
	Type string `json:"type" mapstructure:"type"`
}
