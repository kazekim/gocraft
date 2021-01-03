package gomod

import (
	gcconstants "github.com/kazekim/gocraft/constants"
	"github.com/kazekim/gocraft/filemanager"
	"github.com/kazekim/gocraft/models"
	"text/template"
)

type Template struct {
	PackagePath      string
	RequiredPackages []RequiredPackage
	ReplacedPackage  []ReplacedPackage
}

type RequiredPackage struct {
	Path    string
	Version string
}

type ReplacedPackage struct {
	Path        string
	Version     string
	ReplacePath string
}

func NewTemplate(pkgPath string, exTypes []models.ExternalType) *Template {

	reqPkgs := make([]RequiredPackage, 0)
	repPkgs := make([]ReplacedPackage, 0)
	for _, exType := range exTypes {
		reqPkg := RequiredPackage{
			Path:    exType.GitPath,
			Version: exType.Version,
		}
		reqPkgs = append(reqPkgs, reqPkg)

		if exType.ReplacePath != "" {
			repPkg := ReplacedPackage{
				Path:        exType.GitPath,
				Version:     exType.Version,
				ReplacePath: exType.ReplacePath,
			}
			repPkgs = append(repPkgs, repPkg)
		}
	}
	return &Template{
		PackagePath:      pkgPath,
		RequiredPackages: reqPkgs,
		ReplacedPackage:  repPkgs,
	}
}

func (t *Template) GenerateFile(fileMgr *filemanager.FileManager) error {

	f := fileMgr.NewFile(gcconstants.FileNameGoMod)
	defer f.Close()

	gmTemplate := template.Must(template.New("").Parse(templateStructure))

	err := gmTemplate.Execute(f, t)
	return err
}
