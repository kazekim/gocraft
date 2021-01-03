package models

type CommandConfigSetting struct {
	Package         string           `json:"package" mapstructure:"package"`
	ConfigVariables []ConfigVariable `json:"variables" mapstructure:"variables"`
}
