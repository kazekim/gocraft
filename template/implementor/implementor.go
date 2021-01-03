package implementortemplate

import (
	"github.com/iancoleman/strcase"
	"github.com/kazekim/gocraft/converter/type"
	"github.com/kazekim/gocraft/filemanager"
	"github.com/kazekim/gocraft/models"
	"strings"
	"text/template"
)

type Template struct {
	path               string
	PackageName        string
	InterfaceName      string
	ImplementorName    string
	NewFuncName        string
	Imports            []string
	Attributes         []string
	NewFuncParameters  string
	VariableParameters []string
	Methods            []MethodTemplate
}

type MethodTemplate struct {
	Name             string
	Parameters       string
	ReturnParameters string
	Body             string
	ReturnValues     string
}

func NewTemplate(pkgName, interfaceName string, impl models.Implementor, path string) *Template {

	intfName := strcase.ToCamel(interfaceName)
	implName := strcase.ToLowerCamel(impl.Name) + intfName
	newFuncName := implName

	var imports []string

	attrs, ips := createAttributeTemplates(impl.Attributes)
	imports = appendIfMissing(imports, ips...)

	newFuncParams, variableParams, ips := createParametersTemplates(impl.AllParameters())
	nfpStr := strings.Join(newFuncParams, ",")
	imports = appendIfMissing(imports, ips...)

	//ms := convertMethodTemplate(methods)

	return &Template{
		PackageName:        pkgName,
		InterfaceName:      intfName,
		ImplementorName:    implName,
		NewFuncName:        newFuncName,
		NewFuncParameters:  nfpStr,
		VariableParameters: variableParams,
		Attributes:         attrs,
		Imports:            imports,
		path:               path,
	}
}

func (t *Template) SetNewFuncName(name string) {
	t.NewFuncName = strcase.ToCamel(name)
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

func createParametersTemplates(ps []models.Parameter) ([]string, []string, []string) {
	var params []string
	var vParams []string
	var ips []string

	for _, p := range ps {
		param := p.Name + " "
		vParam := p.Name + ": " + p.Name

		t, ip := typeconverter.FullParameterTypeNameWithImportPath(p)
		param += t
		if ip != "" {
			ips = appendIfMissing(ips, ip)
		}
		params = append(params, param)
		vParams = append(vParams, vParam)
	}
	return params, vParams, ips
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
