package filemanager

import (
	gcutils "github.com/kazekim/gocraft/utils"
	"os"
)

type FileManager struct {
	basePath      string
	fileList      []string
	directoryList []string
}

func New(basePath string) *FileManager {
	return &FileManager{
		basePath:      basePath,
		fileList:      []string{},
		directoryList: []string{},
	}
}

func (fm *FileManager) NewDirectory(directoryName string) {

	err := gcutils.CreateDirectory(fm.basePath, directoryName)
	if err != nil {
		fm.DeleteAllFiles()
		panic(err)
	}
	fm.directoryList = append(fm.directoryList, fm.basePath+directoryName)
}

func (fm *FileManager) NewFile(fileName string) *os.File {
	f, err := os.Create(fm.Path() + "go.mod")
	if err != nil {
		fm.DeleteAllFiles()
		panic(err)
	}
	fm.fileList = append(fm.fileList, fm.basePath+fileName)
	return f
}

func (fm *FileManager) DeleteAllFiles() {

	for _, f := range fm.fileList {
		_ = os.Remove(f)
	}

	for _, d := range fm.directoryList {
		_ = os.RemoveAll(d)
	}
}

func (fm *FileManager) Path() string {
	return fm.basePath
}
