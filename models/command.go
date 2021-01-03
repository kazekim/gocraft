package models

type Command struct {
	Name          string               `json:"name" mapstructure:"name"`
	ConfigSetting CommandConfigSetting `json:"config_setting" mapstructure:"config_setting"`
}
