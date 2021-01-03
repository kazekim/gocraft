package models

type GoCraft struct {
	Prefix           string              `json:"prefix" mapstructure:"prefix"`
	SubPrefix        string              `json:"sub_prefix" mapstructure:"sub_prefix"`
	EnableGoModules  bool                `json:"enable_go_modules" mapstructure:"enable_go_modules"`
	Architecture     ProjectArchitecture `json:"architecture" mapstructure:"architecture"`
	Structure        interface{}         `json:"structure" mapstructure:"structure"`
	ExternalPackages []Package           `json:"external_packages" mapstructure:"external_packages"`
}
