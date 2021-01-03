package models

type Package struct {
	Name               string      `json:"name" mapstructure:"name"`
	Path               string      `json:"path" mapstructure:"path"`
	IsEnableVersioning bool        `json:"is_enable_versioning" mapstructure:"is_enable_versioning"`
	Interfaces         []Interface `json:"interfaces" mapstructure:"interfaces"`
	Enums              []Interface `json:"enums" mapstructure:"enums"`
}
