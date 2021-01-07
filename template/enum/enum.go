package enumtemplate

import (
	"github.com/iancoleman/strcase"
	"github.com/kazekim/gocraft/filemanager"
	"text/template"
)

type Template struct {
	path              string
	PackageName       string
	EnumName          string
	EnumType          string
	EnumTypeCamelCase string
}

func NewTemplate(pkgName, enumName, enumType, path string) *Template {

	eycc := strcase.ToCamel(enumType)
	return &Template{
		PackageName:       pkgName,
		EnumName:          enumName,
		EnumType:          enumType,
		EnumTypeCamelCase: eycc,
		path:              path,
	}
}

func (t *Template) GenerateFile(fileMgr *filemanager.FileManager) {

	f := fileMgr.NewGoFileWithPath(t.path, strcase.ToSnake(t.EnumName))
	defer f.Close()

	gmTemplate := template.Must(template.New("").Parse(templateStructure))

	err := gmTemplate.Execute(f, t)
	if err != nil {
		panic(err)
	}
}
