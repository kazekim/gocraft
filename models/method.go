package models

type Method struct {
	Name              string      `json:"name" mapstructure:"name"`
	Parameters        []Parameter `json:"parameters" mapstructure:"parameters"`
	ReturnParameters  []Parameter `json:"return_parameters" mapstructure:"return_parameters"`
	IsInterfaceMethod bool        `json:"is_interface_method" mapstructure:"is_interface_method"`
	Body              string      `json:"body" mapstructure:"body"`
	BodyTemplate      string      `json:"body_template" mapstructure:"body_template"`
	IsGetter          bool        `json:"is_getter" mapstructure:"is_getter"`
	IsSetter          bool        `json:"is_setter" mapstructure:"is_setter"`
}
