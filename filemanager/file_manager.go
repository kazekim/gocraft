package filemanager

import (
	gcutils "github.com/kazekim/gocraft/utils"
	"os"
	"strings"
)

type FileManager struct {
	basePath       string
	filePaths      []string
	directoryPaths []string
	directoryMap   map[string]string
}

func New(basePath string) *FileManager {
	return &FileManager{
		basePath:       basePath,
		filePaths:      []string{},
		directoryPaths: []string{},
		directoryMap:   map[string]string{},
	}
}

func (fm *FileManager) NewDirectory(names ...string) {

	dirName := strings.Join(names, "/")

	name := names[len(names)-1]

	err := gcutils.CreateDirectory(fm.basePath, dirName)
	if err != nil {
		fm.DeleteAllFiles()
		panic(err)
	}
	fm.directoryPaths = append(fm.directoryPaths, fm.basePath+dirName)
	fm.directoryMap[name] = dirName
}

func (fm *FileManager) NewSubDirectory(parentDir, dirName string) {

	isNotExist, err := gcutils.IsDirectoryNotExist(fm.basePath, parentDir)
	if err != nil {
		panic(err)
	}

	if isNotExist {
		fm.NewDirectory(parentDir)
	}
	fm.NewDirectory(parentDir, dirName)
}

func (fm *FileManager) DeleteAllFiles() {

	for _, f := range fm.filePaths {
		_ = os.Remove(f)
	}

	for _, d := range fm.directoryPaths {
		_ = os.RemoveAll(d)
	}
}

func (fm *FileManager) Path() string {
	return fm.basePath
}
