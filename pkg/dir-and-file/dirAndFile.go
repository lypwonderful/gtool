package dirandfile

import (
	"os"
	"fmt"
)

func isDirExitd(path string) bool{
	_,err := os.Stat(path)
	if os.IsNotExist(err){
		return false
	}
	return true
}
func CreateDir(dirPath string)error{
	if ! isDirExitd(dirPath){
		if err:=os.MkdirAll(dirPath,os.ModePerm);err!=nil{
			fmt.Println(dirPath," is not exist and create it fail")
			return err
		}
		fmt.Println(dirPath," is not exist and create it successed!")
		return nil
	}
	fmt.Println(dirPath," is existed!")
	return nil
}