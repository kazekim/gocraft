package models

type Implementor struct {
	Name       string      `json:"name" mapstructure:"name"`
	Attributes []Attribute `json:"attributes" mapstructure:"attributes"`
	Parameters []Parameter `json:"parameters" mapstructure:"parameters"`
	Methods    []Method    `json:"methods" mapstructure:"methods"`
}

func (i *Implementor) AllParameters() []Parameter {
	var params []Parameter
	for _, a := range i.Attributes {
		if a.IsParameter {
			params = append(params, a.Parameter())
		}
	}
	params = append(params, i.Parameters...)
	return params
}

func (i *Implementor) AllMethods() []Method {

	var ms []Method
	for _, attr := range i.Attributes {
		ms = append(ms, attr.AllMethods()...)
	}
	ms = append(ms, i.Methods...)
	return ms
}
