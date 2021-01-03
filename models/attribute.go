package models

import (
	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
)

type Attribute struct {
	Name              string `json:"name" mapstructure:"name"`
	Type              string `json:"type" mapstructure:"type"`
	Package           string `json:"package" mapstructure:"package"`
	TemplateType      string `json:"template_type" mapstructure:"template_type"`
	IsPointer         bool   `json:"is_pointer" mapstructure:"is_pointer"`
	IsUseCase         bool   `json:"is_use_case" mapstructure:"is_use_case"`
	IsRepository      bool   `json:"is_repository" mapstructure:"is_repository"`
	IsEntity          bool   `json:"is_entity" mapstructure:"is_entity"`
	IsParameter       bool   `json:"is_parameter" mapstructure:"is_parameter"`
	IsAddGetter       bool   `json:"is_add_getter" mapstructure:"is_add_getter"`
	IsAddSetter       bool   `json:"is_add_setter" mapstructure:"is_add_setter"`
	IsInterfaceMethod *bool  `json:"is_interface_method" mapstructure:"is_interface_method"`
}

func (a *Attribute) Parameter() Parameter {
	var p Parameter
	err := mapstructure.Decode(a, &p)
	if err != nil {
		panic(err)
	}
	return p
}

func (a *Attribute) AllMethods() []Method {

	var methods []Method
	var parameter Parameter
	err := mapstructure.Decode(a, &parameter)
	if err != nil {
		panic(err)
	}

	if a.IsAddSetter {
		m := Method{
			Name: "Set" + strcase.ToCamel(a.Name),
			Parameters: []Parameter{
				parameter,
			},
		}
		methods = append(methods, m)
	}

	if a.IsAddGetter {
		m := Method{
			Name: strcase.ToCamel(a.Name),
			ReturnParameters: []Parameter{
				parameter,
			},
		}
		methods = append(methods, m)
	}

	return methods
}
