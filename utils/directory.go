package gcutils

import "os"

func CreateDirectory(path, directoryName string) error {
	fullPath := path + directoryName
	if _, err := os.Stat(fullPath); err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(fullPath, 0755)
			return err
		} else {
			return err
		}
	}

	return nil
}
