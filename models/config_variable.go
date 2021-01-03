package models

type ConfigVariable struct {
	Name    string             `json:"name" mapstructure:"name"`
	Key     string             `json:"key" mapstructure:"key"`
	Default string             `json:"default" mapstructure:"default"`
	Type    ConfigVariableType `json:"type" mapstructure:"type"`
}
