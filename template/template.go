package template

import "github.com/kazekim/gocraft/filemanager"

type Template interface {
	GenerateFile(fileMgr *filemanager.FileManager) error
}
