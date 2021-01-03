package typeconverter

import (
	gcfactory "github.com/kazekim/gocraft/factory"
	"github.com/kazekim/gocraft/models"
)

func FullParameterTypeNameWithImportPath(p models.Parameter) (string, string) {

	t := ""
	ip := ""

	if p.IsPointer {
		t += "*"
	}

	if p.TemplateType != "" {
		et := gcfactory.ExternalTypeFactory().Search(p.TemplateType)
		t += et.PackageName + "." + et.Type
		ip = et.GitPath
	} else {
		if p.Package != "" {
			t += p.Package + "."
		}
		t += p.Type
	}
	return t, ip
}
