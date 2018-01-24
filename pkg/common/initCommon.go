package common

import (
	"fmt"
	"os"
)

func CheckPlatfrom() {
	goroot := os.Getenv("GOROOT")
	if goroot == "" {
		fmt.Println("GOROOT is empty!!!")
		os.Exit(-1)
	}
	fmt.Println("go root:", goroot)
}
