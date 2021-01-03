package cleanarchv1

import (
	"github.com/kazekim/gocraft/architecture/cleanarch/v1/models/v1"
	gcconstants "github.com/kazekim/gocraft/constants"
	"github.com/kazekim/gocraft/filemanager"
	gcgenerator "github.com/kazekim/gocraft/generator"
	"github.com/kazekim/gocraft/models"
	gcstructure "github.com/kazekim/gocraft/structure"
	"github.com/mitchellh/mapstructure"
)

type Structure struct {
	model         cleanarchmodels.CleanArchitecture
	prefix        string
	pkgName       string
	externalTypes []models.ExternalType
}

func NewStructure(m interface{}, prefix, pkgName string, externalTypes []models.ExternalType) gcstructure.Structure {

	var model cleanarchmodels.CleanArchitecture
	err := mapstructure.Decode(m, &model)
	if err != nil {
		panic(err)
	}

	return &Structure{
		model:         model,
		prefix:        prefix,
		pkgName:       pkgName,
		externalTypes: externalTypes,
	}
}

func (s *Structure) Craft(fileMgr *filemanager.FileManager) {

	appPath := gcconstants.DirectoryNameApplciation
	fileMgr.NewDirectory(appPath)
	fileMgr.NewSubDirectory(appPath, gcconstants.DirectoryNameEndPoints)
	fileMgr.NewSubDirectory(appPath, gcconstants.DirectoryNameUseCases)
	fileMgr.NewSubDirectory(appPath, gcconstants.DirectoryNameRepositories)
	fileMgr.NewSubDirectory(appPath, gcconstants.DirectoryNameEntities)
	fileMgr.NewDirectory(gcconstants.DirectoryNamePackages)
	fileMgr.NewDirectory(gcconstants.DirectoryNameInternal)
	fileMgr.NewDirectory(gcconstants.DirectoryNameCommand)

	for _, p := range s.model.Packages {
		g := gcgenerator.NewPackageGenerator(p, s.prefix)
		g.GenerateFile(fileMgr)
	}
}
