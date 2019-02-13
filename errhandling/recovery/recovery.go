package main

import (
	"fmt"
)

func tryRecovery() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occoured.", err)
		} else {
			panic(fmt.Sprintf("I dont know what to do :%v", r))
		}
	}()
	//panic(errors.New("This is a fake error."))
	//b := 0
	//a := 5 / b
	//fmt.Println(a)
	panic(123)
}
func main() {
	tryRecovery()
}
