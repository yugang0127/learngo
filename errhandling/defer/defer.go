package main

import (
	"bufio"
	"errors"
	"fmt"
	"learngo/functional/fib"
	"os"
)

func tryDefer() {
	//fmt.Println("YuWan")
	//fmt.Println("CuMian")
	//defer fmt.Println("123")
	//defer fmt.Println("456")
	//panic("Some Error Happend")
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i >= 30 {
			panic("write too many")
		}
	}

}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	err = errors.New("This is a custom error")
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s,%s,%s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 30; i++ {
		fmt.Fprintln(writer, f())
	}
}
func main() {
	writeFile("fib.txt")
	//tryDefer()

}
