package gcutils

import (
	"go/build"
	"os"
	"strings"
)

func GetCurrentProjectPath() string {

	pth, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	prjPath := strings.Replace(pth, gopath+"/src/", "", 1)
	return prjPath
}
