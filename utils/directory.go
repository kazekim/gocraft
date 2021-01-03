package gcutils

import "os"

func CreateDirectory(path, directoryName string) error {

	isNotExist, err := IsDirectoryNotExist(path, directoryName)
	if err != nil {
		return err
	}

	if isNotExist {
		fullPath := path + directoryName
		err := os.Mkdir(fullPath, 0755)
		return err
	}

	return nil
}

func IsDirectoryNotExist(path, directoryName string) (bool, error) {
	fullPath := path + directoryName
	if _, err := os.Stat(fullPath); err != nil {
		if os.IsNotExist(err) {
			return true, nil
		} else {
			return false, err
		}
	}
	return false, nil
}

func IsPathNotExist(fullPath string) (bool, error) {
	if _, err := os.Stat(fullPath); err != nil {
		if os.IsNotExist(err) {
			return true, nil
		} else {
			return false, err
		}
	}
	return false, nil
}
