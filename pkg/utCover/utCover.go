package utCover

import (
	"fmt"
	"gtool/pkg/flen"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type AllToolT struct {
	FlenT    flen.FlenPathT
	UtCoverT UtCoverInfo
}
type UtCoverInfo struct {
	PkgPath   string
	UtRate    float64
	ExitState string
	ErrInfo   string
}

type UtCoverInter interface {
	doUtExec()
	retResult(outString string, err error)
	calculateUtCover(outString string)
}

func UtCover(cPkgs []string) {
	//utCoverInfo := &UtCoverInfo{}
	utCoverInfo := &AllToolT{}
	for _, v := range cPkgs {
		utCoverInfo.UtCoverT.PkgPath = v
		utCoverInfo.UtCoverT.doUtExec()
		utCoverInfo.FlenT.GenerateFuncLens(v)
		fmt.Printf("%+v\n", utCoverInfo)
	}
}
func (utCoverInfo *UtCoverInfo) calculateUtCover(outString string) {
	//ok  	layered/pkg/ac	0.002s	coverage: 69.2% of statements  ": 69.2%"
	reg := regexp.MustCompile(`(?i:: ).*%`)
	rmRight := strings.TrimRight(reg.FindAllString(outString, -1)[0], "%")
	rmLeft := strings.TrimLeft(rmRight, ": ")
	utCoverInfo.UtRate, _ = strconv.ParseFloat(rmLeft, 64)
}
func (utCoverInfo *UtCoverInfo) retResult(outString string, err error) {
	if err != nil {
		utCoverInfo.ErrInfo = fmt.Sprintf("%s", outString)
		if utCoverInfo.ExitState == "exit status 2" {
			//fmt.Printf("%+v\n", utCoverInfo)
			os.Exit(1)
		}
		utCoverInfo.UtRate = 0
		return
	}
	utCoverInfo.calculateUtCover(outString)
}
func (utCoverInfo *UtCoverInfo) doUtExec() {
	cmdString := fmt.Sprintf("go test -cover  " + utCoverInfo.PkgPath)
	utCmd := exec.Command("bash", "-c", cmdString)
	out, err := utCmd.CombinedOutput()
	utCoverInfo.ExitState = utCmd.ProcessState.String()

	outStr := fmt.Sprintf("%s", string(out))
	utCoverInfo.retResult(outStr, err)
}
