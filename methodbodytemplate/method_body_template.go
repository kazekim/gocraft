package methodbodytemplate

import "errors"

type Template struct {
	Body    string
	Imports []string
}

func Load(templateName string) (*Template, error) {

	switch templateName {
	case TemplateNameReturnFmtString:
		return &templateReturnFmtString, nil
	default:
		return nil, errors.New("template not found")
	}
}
