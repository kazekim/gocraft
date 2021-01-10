package interfacetemplate

import (
	"github.com/iancoleman/strcase"
	"github.com/kazekim/gocraft/filemanager"
	"github.com/kazekim/gocraft/models"
	methodtemplate "github.com/kazekim/gocraft/template/method"
	"text/template"
)

type Template struct {
	path          string
	PackageName   string
	InterfaceName string
	Imports       []string
	Methods       []methodtemplate.MethodTemplate
}

func NewTemplate(pkgName, interfaceName string, methods []models.Method, path string) *Template {

	intfName := strcase.ToCamel(interfaceName)

	ms, ips := methodtemplate.ConvertMethodTemplate(methods)

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
