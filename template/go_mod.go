package gstemplate

import (
	"github.com/kazekim/gocraft/models"
	"os"
	"text/template"
)

type GoModTemplate struct {
	ProjectPath      string
	ProjectName      string
	RequiredPacakges []GoModRequiredPackage
	ReplacedPackage  []GoModReplacedPackage
}

type GoModRequiredPackage struct {
	Path    string
	Version string
}

type GoModReplacedPackage struct {
	Path        string
	Version     string
	ReplacePath string
}

func NewGoModTemplate(prjName, prjPath string, exTypes []models.ExternalType) *GoModTemplate {

	reqPkgs := make([]GoModRequiredPackage, 0)
	repPkgs := make([]GoModReplacedPackage, 0)
	for _, exType := range exTypes {
		reqPkg := GoModRequiredPackage{
			Path:    exType.GitPath,
			Version: exType.Version,
		}
		reqPkgs = append(reqPkgs, reqPkg)

		if exType.ReplacePath != "" {
			repPkg := GoModReplacedPackage{
				Path:        exType.GitPath,
				Version:     exType.Version,
				ReplacePath: exType.ReplacePath,
			}
			repPkgs = append(repPkgs, repPkg)
		}
	}
	return &GoModTemplate{
		ProjectName:      prjName,
		ProjectPath:      prjPath,
		RequiredPacakges: reqPkgs,
		ReplacedPackage:  repPkgs,
	}
}

func (t *GoModTemplate) GenerateFile(path string) error {
	f, err := os.Create(path + "/go.mod")
	if err != nil {
		return err
	}
	defer f.Close()

	gmTemplate := template.Must(template.New("").Parse(goModTemplateStructure))

	err = gmTemplate.Execute(f, t)
	return err
}

const goModTemplateStructure = `module {{.ProjectPath}}{{.ProjectName}}

go 1.15

require (
	{{- range .RequiredPacakges }}
		{{.Path}} {{.Version}}
	{{- end }}
)

replace (
	{{- range .ReplacedPackage }}
		{{.Path}} {{.Version}} => {{.ReplacePath}}
	{{- end }}
)
`
