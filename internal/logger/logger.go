package logger

import "fmt"

func Err(err error) {
	fmt.Println("--- Error --- ", err)
}

func Info(args ...string) {
	fmt.Println("--- INFO --- ", args)
}
