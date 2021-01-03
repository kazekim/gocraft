package models

type Method struct {
	Name              string      `json:"name" mapstructure:"name"`
	Parameters        []Parameter `json:"parameters" mapstructure:"parameters"`
	ReturnParameters  []Parameter `json:"return_parameters" mapstructure:"return_parameters"`
	IsInterfaceMethod bool        `json:"is_interface_method" mapstructure:"is_interface_method"`
}
