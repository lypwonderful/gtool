package main

import (
	"gtool/pkg/cli"
	"gtool/pkg/common"
	"gtool/pkg/dir-and-file"
	"os"
)

var (
	winTestPath   = "F:/tmp/"
	linuxTestPath = "/tmp"
	testPath      string
)

func gtoolInit() {
	testPath = common.CheckPlatfrom()
	switch testPath {
	case "win":
		{
			testPath = winTestPath
		}
	case "linux":
		{
			testPath = linuxTestPath
		}
	default:
		testPath = ""
	}
	if err := dirandfile.CreateDir(testPath); err != nil {
		os.Exit(1)
	}

}
func main() {
	gtoolInit()
	cli.CmdConfig()

}
