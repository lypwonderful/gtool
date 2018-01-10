package utCover

import (
	"fmt"
	"os/exec"
)

type UtCoverInfo struct {
	PkgPath   string
	UtRate    float32
	ExitState int
	ErrInfo   string
}
type UtCoverInter interface {
	doUtExec()
	retResult()
}

func UtCover(cPkgs []string) {
	utCoverInfo := &UtCoverInfo{}
	for _, v := range cPkgs {
		utCoverInfo.PkgPath = v
	}
}
func retResult(outString string, err error) int {
	if err != nil {
		return 0
	}
	fmt.Println("utCover:", outString)
	return 1

}
func (utCoverInfo *UtCoverInfo) doUtExec() {
	cmdString := fmt.Sprintf("go test -cover -v " + utCoverInfo.PkgPath)
	utCmd := exec.Command("bash", "-c", cmdString)
	out, err := utCmd.CombinedOutput()

	outStr := fmt.Sprintf("%s", string(out))
	retResult(outStr, err)
}
