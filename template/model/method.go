package modeltemplate

import (
	"github.com/iancoleman/strcase"
	"github.com/kazekim/gocraft/filemanager"
	"github.com/kazekim/gocraft/models"
	methodtemplate "github.com/kazekim/gocraft/template/method"
	"strings"
	"text/template"
)

type ModelMethodTemplate struct {
	path              string
	PackageName       string
	ModelName         string
	ModelVariableName string
	Method            methodtemplate.MethodTemplate
}

func generateModelMethodTemplates(pkgName, implName string, ms []models.Method, path string, fileMgr *filemanager.FileManager) {

	mts, _ := methodtemplate.ConvertMethodTemplate(ms)

	for _, mt := range mts {

		imt := NewImplementorMethodTemplates(pkgName, implName, mt, path)
		imt.GenerateFile(fileMgr)
	}
}

func NewImplementorMethodTemplates(pkgName, implName string, method methodtemplate.MethodTemplate, path string) *ModelMethodTemplate {

	ivn := string(implName[0])
	return &ModelMethodTemplate{
		path:              path,
		PackageName:       pkgName,
		ModelName:         implName,
		ModelVariableName: ivn,
		Method:            method,
	}
}

func (t *ModelMethodTemplate) GenerateFile(fileMgr *filemanager.FileManager) {

	fileName := strcase.ToSnake(t.ModelName) + "_" + strcase.ToSnake(t.Method.Name)

	f := fileMgr.NewGoFileWithPath(t.path, fileName)
	defer f.Close()

	t.Method.Body = strings.Replace(t.Method.Body, methodtemplate.VariableName, t.ModelVariableName, -1)
	gmTemplate := template.Must(template.New("").Parse(modelMethodTemplateStructure))

	err := gmTemplate.Execute(f, t)
	if err != nil {
		panic(err)
	}
}
