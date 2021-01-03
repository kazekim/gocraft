package models

type Attribute struct {
	Name         string `json:"name" mapstructure:"name"`
	Type         string `json:"type" mapstructure:"type"`
	Package      string `json:"package" mapstructure:"package"`
	TemplateType string `json:"template_type" mapstructure:"template_type"`
	IsPointer    bool   `json:"is_pointer" mapstructure:"is_pointer"`
	IsUseCase    bool   `json:"is_use_case" mapstructure:"is_use_case"`
	IsRepository bool   `json:"is_repository" mapstructure:"is_repository"`
	IsParameter  bool   `json:"is_parameter" mapstructure:"is_parameter"`
	IsEntity     bool   `json:"is_entity" mapstructure:"is_entity"`
}
