package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	a = 2018
 	b = 8
 	c = 30
 	s = "Birthday"
 	f = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitValue()  {
	var a, b int = 3, 4
	var s string = "YuWan"
	fmt.Println(a, b, s)
}

func variableTypeDeduction()  {
	var a, b, s, f = 3, 4, "CuMian", true
	fmt.Println(a, b, s, f)
}

func variableShorter()  {
	a, b, c, s, f := 2018, 8, 30, "Birthday", true
	fmt.Println(a, b, c, s, f)
}

func euler() {
	fmt.Println(cmplx.Abs(3 + 4i))
	fmt.Println(1i * 1i)
	fmt.Println(1i * 1i * 1i)
	fmt.Println(cmplx.Exp(1i * math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i * math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i * 3 / 2 * math.Pi))
}

func triangle()  {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	return int(math.Sqrt(float64(a * a + b * b)))
}

func consts()  {
	const filename = "aaa.txt"
	const a int = 3
	const b int = 4
	const (
		d = 3
		e = 4
	)
	var c, f int
	c = int(math.Sqrt(float64(a * a + b * b)))
	f = int(math.Sqrt(d * d + e * e))


	fmt.Println(filename, a, b, c, f)
}

func enums()  {
const (
	cpp = iota * 3 + 1
	java
	php
	python
	golang
)

const (
	b = 1 << (10 * iota)
	kb
	mb
	gb
	tb
	pb
)

	fmt.Println(cpp, java, php, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello World!")
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(a, b, c, s, f)
	euler()
	triangle()
	consts()
	enums()
}
