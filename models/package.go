package models

type Package struct {
	Name               string      `json:"name" mapstructure:"name"`
	Type               PackageType `json:"type" mapstructure:"type"`
	CustomPath         string      `json:"custom_path" mapstructure:"custom_path"`
	IsAddPrefix        bool        `json:"is_add_prefix" mapstructure:"is_add_prefix"`
	IsEnableVersioning bool        `json:"is_enable_versioning" mapstructure:"is_enable_versioning"`
	Interfaces         []Interface `json:"interfaces" mapstructure:"interfaces"`
	Enums              []Enum      `json:"enums" mapstructure:"enums"`
}

type PackageType string

const (
	PackageTypePackage  PackageType = "pkg"
	PackageTypeInternal PackageType = "internal"
	PackageTypeCustom   PackageType = "custom"
)

func (t PackageType) String() string {
	return string(t)
}
