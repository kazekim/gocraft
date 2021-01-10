package models

type Model struct {
	Name           string      `json:"name" mapstructure:"name"`
	Attributes     []Attribute `json:"attributes" mapstructure:"attributes"`
	Parameters     []Parameter `json:"parameters" mapstructure:"parameters"`
	Methods        []Method    `json:"methods" mapstructure:"methods"`
	HasConstructor bool        `json:"has_constructor" mapstructure:"has_constructor"`
}

func (m *Model) AllParameters() []Parameter {
	var params []Parameter
	for _, a := range m.Attributes {
		if a.IsParameter {
			params = append(params, a.Parameter())
		}
	}
	params = append(params, m.Parameters...)
	return params
}

func (m *Model) AllMethods() []Method {

	var ms []Method
	for _, attr := range m.Attributes {
		ms = append(ms, attr.AllMethods()...)
	}
	ms = append(ms, m.Methods...)
	return ms
}
