package main

import "fmt"

func main() {
	m := map[string]string {
		"name":	"YuWan",
		"Birthday":	"20180830",
		"From":	"ChongQing",
		"Love":	"Piano",
	}

	m2 := make(map[string]int)
	var m3 map[string]int

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map:")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values:")
	name, ok := m["name"]
	fmt.Println(name, ok)

	if namm, ok := m["namm"]; ok {
		fmt.Println(namm, ok)
	} else {
		fmt.Println("key namm does not found")
	}

	fmt.Println("Delete keys")
	birthDay, ok := m["Birthday"]
	fmt.Printf("m[%q] before delete %q, %v\n ",
		"birthday", birthDay, ok)
	delete(m,"Birthday")
	birthDay, ok = m["Birthday"]
	fmt.Printf("m[%q] after delete %q, %v\n ",
		"birthday", birthDay, ok)
}
