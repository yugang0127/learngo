package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation %s", op)
	}
}

func div(a, b int) (q, r int)  {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int  {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling func %s with args" +
		"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func sumArgs(values ...int) int  {
	sum := 0
	for i := range values {
		sum += values[i]
	}

	return sum
}

func swapWithPointer(a, b *int) {
	*b, *a = *a, *b
}

func swap(a, b int) (int, int) {
	return b, a
}


func main() {
	if result, err := eval(3, 4, "*"); err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(99, 18)
	fmt.Printf("99 div 18 is %d mod %d\n", q, r)

	fmt.Println("Power(3, 4) is:", apply(
		func(a, b int) int {
			return int(math.Pow(
				float64(a), float64(b)))
		}, 3, 4))
	fmt.Println("sumArgs(1, 2, 3...5) is ", sumArgs(1, 2, 3, 4, 5))

	a, b := 10, 100
	a, b = swap(a, b)
	fmt.Println("a, b after swap is :", a, b)

	c, d := 1987, 1988
	swapWithPointer(&c, &d)
	fmt.Println("c, d after swap is :", c, d)


	var ai int = 2
	var pa *int = &ai
	*pa = 3
	fmt.Println(ai)
}
