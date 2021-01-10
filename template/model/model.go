package modeltemplate

import (
	"github.com/iancoleman/strcase"
	typeconverter "github.com/kazekim/gocraft/converter/type"
	"github.com/kazekim/gocraft/filemanager"
	"github.com/kazekim/gocraft/models"
	"strings"
	"text/template"
)

type Template struct {
	path               string
	PackageName        string
	ModelName          string
	NewFuncName        string
	Imports            []string
	Attributes         []string
	NewFuncParameters  string
	VariableParameters []string
	Methods            []models.Method
}

func NewTemplate(pkgName string, m models.Model, path string) *Template {

	modelName := strcase.ToCamel(m.Name)
	newFuncName := modelName

	var imports []string

	attrs, ips := createAttributeTemplates(m.Attributes)
	imports = appendIfMissing(imports, ips...)

	newFuncParams, variableParams, ips := createParametersTemplates(m.AllParameters())
	nfpStr := strings.Join(newFuncParams, ", ")
	imports = appendIfMissing(imports, ips...)

	return &Template{
		PackageName:        pkgName,
		ModelName:          modelName,
		NewFuncName:        newFuncName,
		NewFuncParameters:  nfpStr,
		VariableParameters: variableParams,
		Attributes:         attrs,
		Imports:            imports,
		path:               path,
		Methods:            m.Methods,
	}
}

func (t *Template) GenerateFile(fileMgr *filemanager.FileManager) {

	f := fileMgr.NewGoFileWithPath(t.path, strcase.ToSnake(t.ModelName))
	defer f.Close()

	gmTemplate := template.Must(template.New("").Parse(templateStructure))

	err := gmTemplate.Execute(f, t)
	if err != nil {
		panic(err)
	}

	generateModelMethodTemplates(t.PackageName, t.ModelName, t.Methods, t.path, fileMgr)
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
		param := strcase.ToLowerCamel(p.Name) + " "
		vParam := p.Name + ": " + strcase.ToLowerCamel(p.Name)

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
		isDuplicate := false
		for _, ele := range slice {
			if ele == val {
				isDuplicate = true
			}
		}
		if !isDuplicate {
			slice = append(slice, val)
		}
	}

	return slice
}
