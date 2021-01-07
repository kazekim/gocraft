package gcgenerator

import "github.com/kazekim/gocraft/filemanager"

type Generator interface {
	GenerateFile(fileMgr *filemanager.FileManager)
}
