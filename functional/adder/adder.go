package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}
func main() {
	a := adder()
	for i := 1; i <= 100; i++ {
		fmt.Printf("0 + 1 + ... %d = %d\n", i, a(i))
	}
}
