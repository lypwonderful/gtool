package main

import (
	"fmt"
	"gtool/pkg/cli"
	"gtool/pkg/common"
	"gtool/pkg/dir-and-file"
	"gtool/pkg/utCover"
	"os"
)

var testPath = "F:/tmp/"

func gtoolInit() {
	if err := dirandfile.CreateDir(testPath); err != nil {
		os.Exit(1)
	}
	fmt.Println("checkPath:", common.ReadArgs().CheckPkgs)
	utCover.UtCover(common.ReadArgs().CheckPkgs)
}
func main() {

	cli.CmdConfig()
	gtoolInit()
}
