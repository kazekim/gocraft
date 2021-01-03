package interfacetemplate

import (
	"github.com/iancoleman/strcase"
	"github.com/kazekim/gocraft/filemanager"
	"github.com/kazekim/gocraft/models"
	"text/template"
)

type Template struct {
	path          string
	PackageName   string
	InterfaceName string
	Methods       []MethodTemplate
}

type MethodTemplate struct {
	Name             string
	Parameters       string
	ReturnParameters string
}

func NewTemplate(pkgName, interfaceName, path string, methods []models.Method) *Template {

	intfName := strcase.ToCamel(interfaceName)

	ms := convertMethodTemplate(methods)

	return &Template{
		PackageName:   pkgName,
		InterfaceName: intfName,
		Methods:       ms,
		path:          path,
	}
}

func (t *Template) GenerateFile(fileMgr *filemanager.FileManager) {

	f := fileMgr.NewGoFileWithPath(t.path, strcase.ToSnake(t.InterfaceName))
	defer f.Close()

	gmTemplate := template.Must(template.New("").Parse(templateStructure))

	err := gmTemplate.Execute(f, t)
	if err != nil {
		panic(err)
	}
}

func convertMethodTemplate(ms []models.Method) []MethodTemplate {

	var vs []MethodTemplate
	for _, m := range ms {

		params := ""
		for i, ps := range m.Parameters {
			if i > 0 {
				params = params + ", "
			}
			params = params + ps.Name + " "
			if ps.TemplateType != "" {

			} else {
				if ps.Package != "" {
					params = params + ps.Package + "."
				}
				params = params + ps.Type
			}
		}

		retParams := ""
		if len(m.ReturnParameters) > 1 {
			retParams += "("
		}
		for i, ps := range m.ReturnParameters {
			if i > 0 {
				retParams = retParams + ", "
			}
			if ps.TemplateType != "" {

			} else {
				if ps.Package != "" {
					retParams = retParams + " " + ps.Package + "."
				}
				retParams = retParams + ps.Type
			}
		}
		if len(m.ReturnParameters) > 1 {
			retParams += ")"
		}

		v := MethodTemplate{
			Name:             m.Name,
			Parameters:       params,
			ReturnParameters: retParams,
		}
		vs = append(vs, v)
	}
	return vs
}
