package utCover

import (
	"fmt"
	"os/exec"
)

func UtCover(cPkgs []string) {

	for _, v := range cPkgs {
		cmdString := fmt.Sprintf("go test -cover " + v)
		out, _ := exec.Command("bash", "-c", cmdString).CombinedOutput()
		fmt.Println("utout:", string(out))
	}
}
