package main

import (
	"fmt"
	"regexp"
)

func main() {
	reg := regexp.MustCompile(`(?i:: ).*%`)
	fmt.Printf("%q\n", reg.FindAllString("ok  	layered/pkg/ac	0.002s	coverage: 69.2% of statements", -1)[0])
}
