package gcgenerator

import (
	gcconstants "github.com/kazekim/gocraft/constants"
	"github.com/kazekim/gocraft/filemanager"
	"github.com/kazekim/gocraft/models"
	enumtemplate "github.com/kazekim/gocraft/template/enum"
	gcutils "github.com/kazekim/gocraft/utils"
)

type packageGenerator struct {
	pkg    models.Package
	prefix string
}

func NewPackageGenerator(pkg models.Package, prefix string) Generator {
	return &packageGenerator{
		pkg:    pkg,
		prefix: prefix,
	}
}

func (g *packageGenerator) PackageName() string {
	return g.prefix + g.pkg.Name
}

func (g *packageGenerator) GenerateFile(fileMgr *filemanager.FileManager) {

	pd := ""
	switch g.pkg.Type {
	case models.PackageTypePackage:
		pd = gcconstants.DirectoryNamePackages
	case models.PackageTypeInternal:
		pd = gcconstants.DirectoryNameInternal
	case models.PackageTypeCustom:
		if g.pkg.CustomPath == "" {
			panic("custom path not defined")
		}
		pd = g.pkg.CustomPath
	default:
		panic("package type not defined")
	}
	path := pd + "/" + g.PackageName()
	isNotExist, err := gcutils.IsPathNotExist(fileMgr.Path() + path)
	if err != nil {
		panic(err)
	}
	if isNotExist {
		fileMgr.NewSubDirectory(pd, g.PackageName())
	}
	for _, enum := range g.pkg.Enums {
		et := enumtemplate.NewTemplate(g.PackageName(), enum.Name, enum.Type, path)
		et.GenerateFile(fileMgr)
	}
}
