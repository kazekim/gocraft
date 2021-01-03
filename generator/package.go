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
	pkgName := g.pkg.Name
	if g.pkg.IsAddPrefix {
		pkgName = g.prefix + pkgName
	}
	return pkgName
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

	path := g.selectPathOrCreateDirectoryIfNotExist(fileMgr, pd, g.PackageName())
	if g.pkg.IsEnableVersioning {
		path = g.selectPathOrCreateDirectoryIfNotExist(fileMgr, path, gcconstants.Version1)
	}

	for _, enum := range g.pkg.Enums {
		et := enumtemplate.NewTemplate(g.PackageName(), enum.Name, enum.Type, path)
		et.GenerateFile(fileMgr)
	}
}

func (g *packageGenerator) selectPathOrCreateDirectoryIfNotExist(fileMgr *filemanager.FileManager, path, directory string) string {
	fullPath := path + "/" + directory
	isNotExist, err := gcutils.IsPathNotExist(fileMgr.Path() + fullPath)
	if err != nil {
		panic(err)
	}
	if isNotExist {
		fileMgr.NewSubDirectory(path, directory)
	}
	return fullPath
}
