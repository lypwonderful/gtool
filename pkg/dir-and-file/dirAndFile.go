package dirandfile

import (
	"fmt"
	"os"
)

func IsDirExitd(path string) bool {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		return true
	}
	return false
}
func CreateDir(dirPath string) error {
	if !IsDirExitd(dirPath) {
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			fmt.Println(dirPath, " is not exist and create it fail")
			return err
		}
		fmt.Println(dirPath, " is not exist and create it successed!")
		return nil
	}
	fmt.Println(dirPath, " is existed!")
	return nil
}
