package implementortemplate

import (
	"github.com/iancoleman/strcase"
	"github.com/kazekim/gocraft/converter/type"
	"github.com/kazekim/gocraft/filemanager"
	"github.com/kazekim/gocraft/models"
	"text/template"
)

type Template struct {
	path            string
	PackageName     string
	ImplementorName string
	Imports         []string
	Attributes      []string
	Methods         []MethodTemplate
}

type MethodTemplate struct {
	Name             string
	Parameters       string
	ReturnParameters string
	Body             string
	ReturnValues     string
}

func NewTemplate(pkgName, interfaceName string, impl models.Implementor, path string) *Template {

	implName := strcase.ToCamel(impl.Name) + strcase.ToCamel(interfaceName)

	var imports []string

	attrs, ips := createAttributeTemplates(impl.Attributes)
	imports = appendIfMissing(imports, ips...)
	//ms := convertMethodTemplate(methods)

	return &Template{
		PackageName:     pkgName,
		ImplementorName: implName,
		Attributes:      attrs,
		Imports:         imports,
		path:            path,
	}
}

func (t *Template) GenerateFile(fileMgr *filemanager.FileManager) {

	f := fileMgr.NewGoFileWithPath(t.path, strcase.ToSnake(t.ImplementorName))
	defer f.Close()

	gmTemplate := template.Must(template.New("").Parse(templateStructure))

	err := gmTemplate.Execute(f, t)
	if err != nil {
		panic(err)
	}
}

func createAttributeTemplates(as []models.Attribute) ([]string, []string) {

	var attrs []string
	var ips []string

	for _, a := range as {
		attr := a.Name + " "

		t, ip := typeconverter.FullAttributeTypeNameWithImportPath(a)
		attr += t
		if ip != "" {
			ips = appendIfMissing(ips, ip)
		}
		attrs = append(attrs, attr)
	}
	return attrs, ips
}

func appendIfMissing(slice []string, vals ...string) []string {
	for _, val := range vals {
		for _, ele := range slice {
			if ele == val {
				continue
			}
		}
		slice = append(slice, val)
	}

	return slice
}
