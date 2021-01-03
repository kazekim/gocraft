package interfacetemplate

import (
	"github.com/iancoleman/strcase"
	"github.com/kazekim/gocraft/filemanager"
	"text/template"
)

type Template struct {
	path          string
	PackageName   string
	InterfaceName string
}

func NewTemplate(pkgName, interfaceName, path string) *Template {

	intfName := strcase.ToCamel(interfaceName)
	return &Template{
		PackageName:   pkgName,
		InterfaceName: intfName,
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
