package cleanarch

import (
	"github.com/kazekim/gocraft/cleanarch/models/v1"
	gcconstants "github.com/kazekim/gocraft/constants"
	"github.com/kazekim/gocraft/filemanager"
	"github.com/kazekim/gocraft/models"
	gcstructure "github.com/kazekim/gocraft/structure"
)

type CleanArchitectureStructureV1 struct {
	model         cleanarchmodels.CleanArchitecture
	prefix        string
	pkgName       string
	externalTypes []models.ExternalType
}

func NewStructureV1(model cleanarchmodels.CleanArchitecture, prefix, pkgName string, externalTypes []models.ExternalType) gcstructure.Structure {
	return &CleanArchitectureStructureV1{
		model:         model,
		prefix:        prefix,
		pkgName:       pkgName,
		externalTypes: externalTypes,
	}
}

func (s *CleanArchitectureStructureV1) Craft(fileMgr *filemanager.FileManager) error {

	appPath := gcconstants.DirectoryNameApplciation
	fileMgr.NewDirectory(appPath)
	fileMgr.NewSubDirectory(appPath, gcconstants.DirectoryNameEndPoints)
	fileMgr.NewSubDirectory(appPath, gcconstants.DirectoryNameUseCases)
	fileMgr.NewSubDirectory(appPath, gcconstants.DirectoryNameRepositories)
	fileMgr.NewSubDirectory(appPath, gcconstants.DirectoryNameEntities)
	fileMgr.NewDirectory(gcconstants.DirectoryNamePackages)
	fileMgr.NewDirectory(gcconstants.DirectoryNameInternal)
	fileMgr.NewDirectory(gcconstants.DirectoryNameCommand)

	return nil
}
