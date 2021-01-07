package crafter

import "github.com/kazekim/gocraft/filemanager"

type Crafter interface {
	Craft() error
}

type defaultCrafter struct {
	cfgName string
	cfgPath string
	fileMgr *filemanager.FileManager
}

func New(cfgName, cfgPath string) Crafter {

	fileMgr := filemanager.New(cfgPath)

	return &defaultCrafter{
		cfgName: cfgName,
		cfgPath: cfgPath,
		fileMgr: fileMgr,
	}
}
