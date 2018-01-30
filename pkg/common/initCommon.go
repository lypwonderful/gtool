package common

import (
	"fmt"
	"gtool/pkg/dir-and-file"
	"gtool/pkg/flen"
	"gtool/pkg/utCover"
	"os"
	"path"
	"strings"
)

type AllToolT struct {
	FlenT    flen.FuncLenT
	UtCoverT utCover.UtCoverInfo
}

func CheckPlatfrom() string {
	goroot := os.Getenv("GOROOT")
	if goroot == "" {
		fmt.Println("GOROOT is empty!!!")
		os.Exit(-1)
	}
	return "linux"
}
func getGOPATH() string {
	GOPATH := os.Getenv("GOPATH")
	if GOPATH == "" {
		fmt.Println("GOPATH is empty,Please Set it.")
		os.Exit(-1)
	}
	return GOPATH
}
func IsPkgInGOPATH(pkgPath string) bool {
	gopath := getGOPATH()
	goPaths := strings.Split(gopath, ";")
	for _, v := range goPaths {
		v = path.Join(v, "src", pkgPath)
		if dirandfile.IsDirExitd(v) {
			return true
		}
	}
	return false
}
