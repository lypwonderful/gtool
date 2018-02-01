package common

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type CliArgs struct {
	pkgPath    string
	filterPkgs []string
	CheckPkgs  []string
	len        int
}

func ReadArgs() CliArgs {
	flag.Parse()
	args := flag.Args()
	cliArgs := &CliArgs{}

	cliArgs.len = len(args)
	if cliArgs.len <= 1 {
		flag.Usage()
		os.Exit(-1)
	}

	cliArgs.pkgPath = args[1]
	if IsPkgInGOPATH(cliArgs.pkgPath) {
		fmt.Println("pkg is not in GOPATH! ", cliArgs.pkgPath)
		os.Exit(1)
	}
	if cliArgs.len == 2 {
		cliArgs.filterPkgs = append(cliArgs.filterPkgs, "")
		cliArgs.findChildPkgs()
		return *cliArgs
	}

	cliArgs.filterPkgs = append(cliArgs.filterPkgs, args[2:]...)
	cliArgs.getNewfilterPkgs()
	cliArgs.findChildPkgs()
	return *cliArgs

}
func (cliArgs *CliArgs) getNewfilterPkgs() {
	for k, v := range cliArgs.filterPkgs {
		if strings.HasSuffix(v, "/") {
			cliArgs.filterPkgs[k] = strings.TrimRight(v, "/")
		}
	}
}
func isFilterPkg(src string, dest []string) bool {
	if strings.HasSuffix(src, "/") {
		src = strings.TrimRight(src, "/")
	}
	for _, v := range dest {
		if strings.Contains(src, v) {
			return false
		}
	}
	return true

}
func (cliArgs *CliArgs) findChildPkgs() {
	filepath.Walk(cliArgs.pkgPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".git") {
			return nil
		}
		if cliArgs.len != 2 {
			if info.IsDir() && isFilterPkg(path, cliArgs.filterPkgs) {
				cliArgs.CheckPkgs = append(cliArgs.CheckPkgs, path)
			}
			return nil
		}
		if info.IsDir() {
			cliArgs.CheckPkgs = append(cliArgs.CheckPkgs, path)
		}
		return nil

	})
}
