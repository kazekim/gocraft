package typeconverter

import (
	gcfactory "github.com/kazekim/gocraft/factory"
	"github.com/kazekim/gocraft/models"
)

func FullAttributeTypeNameWithImportPath(a models.Attribute) (string, string) {

	t := ""
	ip := ""

	if a.IsPointer {
		t += "*"
	}

	if a.TemplateType != "" {
		et := gcfactory.ExternalTypeFactory().Search(a.TemplateType)
		t += et.PackageName + "." + et.Type
		ip = et.GitPath
	} else {
		if a.Package != "" {
			t += a.Package + "."
		}
		t += a.Type
	}
	return t, ip
}
