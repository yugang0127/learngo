package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string  {
	g := ""
	switch  {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 70:
		g = "D"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}

	return g;
}

func main() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename);err != nil {
		fmt.Println("Cannot print file contents", err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(58),
		grade(68),
		grade(79),
		grade(88),
		grade(99),
		grade(108),
	)
}
