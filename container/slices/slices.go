package main

import "fmt"

func updateSlices(arr []int) {
	arr[0] = 2018
}

func main() {
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(arr)
	fmt.Println("arr[2:6] :", arr[2:6])
	fmt.Println("arr[2:] :", arr[2:])
	fmt.Println("arr[:6] :", arr[:6])
	fmt.Println("arr[:] :", arr[:])

	s1 := arr[2:6]
	fmt.Println("s1:", s1)
	s2 := arr[:6]
	fmt.Println("s2:", s2)

	fmt.Println("After updateSlices s1:")
	updateSlices(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlices s2:")
	updateSlices(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice:")
	fmt.Println("s2:", s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending Slice:")
	arr[0], arr[2] = 0, 2
	fmt.Println("arr:", arr)
	s1 = arr[2:6]
	s2 = s1[3:5]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",
		s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",
		s2, len(s2), cap(s2))

	fmt.Printf("arr=%v, len(arr)=%d, cap(arr)=%d\n",
		arr, len(arr), cap(arr))
	s3 := append(s2, 9)
	s4 := append(s3, 10)
	s5 := append(s4, 11)
	fmt.Println("s3, s4, s5 = ", s3, s4, s5)
	fmt.Println("arr = ", arr)
}
