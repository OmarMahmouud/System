package common

import "fmt"

func CheckError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
