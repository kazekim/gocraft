package models

import (
	"strings"
)

type GoCraft struct {
	Prefix            string              `json:"prefix" mapstructure:"prefix"`
	PackageName       string              `json:"package_name" mapstructure:"package_name"`
	PackagePath       string              `json:"package_path" mapstructure:"package_path"`
	IsEnableGoModules bool                `json:"is_enable_go_modules" mapstructure:"is_enable_go_modules"`
	Architecture      ProjectArchitecture `json:"architecture" mapstructure:"architecture"`
	Structure         interface{}         `json:"structure" mapstructure:"structure"`
	ExternalTypes     []ExternalType      `json:"external_types" mapstructure:"external_types"`
}

func (gc *GoCraft) UpperCasePrefix() string {
	return strings.ToUpper(gc.Prefix)
}
