package gstemplate

import (
	gcconstants "github.com/kazekim/gocraft/constants"
	"github.com/kazekim/gocraft/filemanager"
	"github.com/kazekim/gocraft/models"
	"text/template"
)

type GoModTemplate struct {
	PackagePath      string
	RequiredPackages []GoModRequiredPackage
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

func NewGoModTemplate(pkgPath string, exTypes []models.ExternalType) *GoModTemplate {

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
		PackagePath:      pkgPath,
		RequiredPackages: reqPkgs,
		ReplacedPackage:  repPkgs,
	}
}

func (t *GoModTemplate) GenerateFile(fileMgr *filemanager.FileManager) error {

	f := fileMgr.NewFile(gcconstants.FileNameGoMod)
	defer f.Close()

	gmTemplate := template.Must(template.New("").Parse(goModTemplateStructure))

	err := gmTemplate.Execute(f, t)
	return err
}

const goModTemplateStructure = `module {{.PackagePath}}

go 1.15

require (
	{{- range .RequiredPackages }}
		{{.Path}} {{.Version}}
	{{- end }}
)

replace (
	{{- range .ReplacedPackage }}
		{{.Path}} {{.Version}} => {{.ReplacePath}}
	{{- end }}
)
`
