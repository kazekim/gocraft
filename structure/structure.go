package gcstructure

import "github.com/kazekim/gocraft/filemanager"

type Structure interface {
	Craft(fileMgr *filemanager.FileManager)
}
