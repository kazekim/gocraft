package models

type GoCraft struct {
	Prefix            string              `json:"prefix" mapstructure:"prefix"`
	SubPrefix         string              `json:"sub_prefix" mapstructure:"sub_prefix"`
	IsEnableGoModules bool                `json:"is_enable_go_modules" mapstructure:"is_enable_go_modules"`
	Architecture      ProjectArchitecture `json:"architecture" mapstructure:"architecture"`
	Structure         interface{}         `json:"structure" mapstructure:"structure"`
	ExternalTypes     []ExternalType      `json:"external_types" mapstructure:"external_types"`
}
