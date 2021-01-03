package filemanager

import "os"

func (fm *FileManager) NewFile(fileName string) *os.File {
	return fm.NewFileWithPath("", fileName)
}

func (fm *FileManager) NewGoFile(fileName string) *os.File {
	return fm.NewFileWithPath("", fileName+".go")
}

func (fm *FileManager) NewFileWithPath(path, fileName string) *os.File {

	if path != "" {
		fileName = path + "/" + fileName
	}

	f, err := os.Create(fm.Path() + fileName)
	if err != nil {
		fm.DeleteAllFiles()
		panic(err)
	}
	fm.filePaths = append(fm.filePaths, fm.basePath+fileName)
	return f
}

func (fm *FileManager) NewGoFileWithPath(path, fileName string) *os.File {
	return fm.NewFileWithPath(path, fileName+".go")
}
