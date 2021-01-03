package models

type ExternalType struct {
	Name        string `json:"name" mapstructure:"name"`
	Type        string `json:"type" mapstructure:"type"`
	PackageName string `json:"package_name" mapstructure:"package_name"`
	GitPath     string `json:"git_path" mapstructure:"git_path"`
}
