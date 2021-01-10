package modelgenerator

import (
	gcconstants "github.com/kazekim/gocraft/constants"
	"github.com/kazekim/gocraft/filemanager"
	"github.com/kazekim/gocraft/generator"
	"github.com/kazekim/gocraft/models"
	modeltemplate "github.com/kazekim/gocraft/template/model"
)

const (
	PackageName = "models"
)

type modelGenerator struct {
	mdls        []models.Model
	prefix      string
	isAddPrefix bool
}

func New(mdls []models.Model, prefix string, isAddPrefix bool) gcgenerator.Generator {
	return &modelGenerator{
		mdls:        mdls,
		prefix:      prefix,
		isAddPrefix: isAddPrefix,
	}
}

func (g *modelGenerator) PackageName() string {
	pkgName := PackageName
	if g.isAddPrefix {
		pkgName = g.prefix + pkgName
	}
	return pkgName
}

func (g *modelGenerator) GenerateFile(fileMgr *filemanager.FileManager) {

	path := gcconstants.DirectoryNameApplciation + "/" + gcconstants.DirectoryNameModels
	for _, m := range g.mdls {

		mt := modeltemplate.NewTemplate(g.PackageName(), m, path)
		mt.GenerateFile(fileMgr)
	}
}
