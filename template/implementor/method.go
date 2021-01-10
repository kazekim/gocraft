package implementortemplate

import (
	"github.com/iancoleman/strcase"
	"github.com/kazekim/gocraft/filemanager"
	"github.com/kazekim/gocraft/models"
	methodtemplate "github.com/kazekim/gocraft/template/method"
	"strings"
	"text/template"
)

type ImplementorMethodTemplate struct {
	path                    string
	PackageName             string
	ImplementorName         string
	ImplementorVariableName string
	Method                  methodtemplate.MethodTemplate
}

func generateImplementorMethodTemplates(pkgName, implName string, ms []models.Method, path string, fileMgr *filemanager.FileManager) {

	mts, _ := methodtemplate.ConvertMethodTemplate(ms)

	for _, mt := range mts {

		imt := NewImplementorMethodTemplates(pkgName, implName, mt, path)
		imt.GenerateFile(fileMgr)
	}
}

func NewImplementorMethodTemplates(pkgName, implName string, method methodtemplate.MethodTemplate, path string) *ImplementorMethodTemplate {

	ivn := string(implName[0])
	return &ImplementorMethodTemplate{
		path:                    path,
		PackageName:             pkgName,
		ImplementorName:         implName,
		ImplementorVariableName: ivn,
		Method:                  method,
	}
}

func (t *ImplementorMethodTemplate) GenerateFile(fileMgr *filemanager.FileManager) {

	fileName := strcase.ToSnake(t.ImplementorName) + "_" + strcase.ToSnake(t.Method.Name)

	f := fileMgr.NewGoFileWithPath(t.path, fileName)
	defer f.Close()

	t.Method.Body = strings.Replace(t.Method.Body, methodtemplate.VariableName, t.ImplementorVariableName, -1)
	gmTemplate := template.Must(template.New("").Parse(implementorMethodTemplateStructure))

	err := gmTemplate.Execute(f, t)
	if err != nil {
		panic(err)
	}
}
