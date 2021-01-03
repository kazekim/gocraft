package models

type Parameter struct {
	Name         string `json:"name" mapstructure:"name"`
	Type         string `json:"type" mapstructure:"type"`
	Package      string `json:"package" mapstructure:"package"`
	TemplateType string `json:"template_type" mapstructure:"template_type"`
	IsUseCase    bool   `json:"is_use_case" mapstructure:"is_use_case"`
	IsRepository bool   `json:"is_repository" mapstructure:"is_repository"`
	IsEntity     bool   `json:"is_entity" mapstructure:"is_entity"`
}
