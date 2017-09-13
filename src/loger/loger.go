package loger

import "fmt"

func CheckError(err error) {
	if err != nil {
		Error(err)
	}
}

func CheckLog(err error) {
	if err != nil {
		Log(err)
	}
}

func Error(err error) {
	panic(err)
}

func Log(err error) {
	fmt.Println(err)
}