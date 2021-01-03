package models

type Interface struct {
	Name               string        `json:"name" mapstructure:"name"`
	DefaultImplementor *Implementor  `json:"default_implementor" mapstructure:"default_implementor"`
	Implementors       []Implementor `json:"implementors" mapstructure:"implementors"`
}

func (i *Interface) AllMethods() []Method {

	var ms []Method
	if impl := i.DefaultImplementor; impl != nil {
		ms = impl.AllMethods()
	}

	return ms
}
