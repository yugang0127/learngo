package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func sum(n int) int  {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}

func convertToBin(n int) string  {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string)  {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever()  {
	for {
		fmt.Println("Hello")
	}
}

func main() {
	fmt.Println(
		sum(100),
		sum(1000),
		sum(10000),
		convertToBin(13),
		convertToBin(128),
		convertToBin(8322),
		convertToBin(20180830),
	)
	fmt.Println()

	printFile("basic/branch/abc.txt")
	s := `Hello
	123
	good"
	Abcd'
	Good Good Study!
	Day Day Up!
	`
	printFileContents(strings.NewReader(s))
	//forever()
}
