package interfacetemplate

import (
	"github.com/iancoleman/strcase"
	typeconverter "github.com/kazekim/gocraft/converter/type"
	"github.com/kazekim/gocraft/filemanager"
	"github.com/kazekim/gocraft/models"
	"text/template"
)

type Template struct {
	path          string
	PackageName   string
	InterfaceName string
	Imports       []string
	Methods       []MethodTemplate
}

type MethodTemplate struct {
	Name             string
	Parameters       string
	ReturnParameters string
}

func NewTemplate(pkgName, interfaceName string, methods []models.Method, path string) *Template {

	intfName := strcase.ToCamel(interfaceName)

	ms, ips := generateMethodTemplate(methods)

	return &Template{
		PackageName:   pkgName,
		InterfaceName: intfName,
		Imports:       ips,
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

func generateMethodTemplate(ms []models.Method) ([]MethodTemplate, []string) {

	var vs []MethodTemplate
	var ips []string
	for _, m := range ms {

		params := ""
		for i, p := range m.Parameters {
			if i > 0 {
				params = params + ", "
			}
			params = params + p.Name + " "

			if p.IsPointer {
				params += "*"
			}

			t, ip := typeconverter.FullParameterTypeNameWithImportPath(p)
			params += t
			if ip != "" {
				ips = appendIfMissing(ips, ip)
			}
		}

		retParams := ""
		if len(m.ReturnParameters) > 1 {
			retParams += "("
		}
		for i, p := range m.ReturnParameters {
			if i > 0 {
				retParams = retParams + ", "
			}

			t, ip := typeconverter.FullParameterTypeNameWithImportPath(p)
			retParams += t
			if ip != "" {
				ips = appendIfMissing(ips, ip)
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
	return vs, ips
}

func appendIfMissing(slice []string, val string) []string {
	for _, ele := range slice {
		if ele == val {
			return slice
		}
	}
	return append(slice, val)
}
