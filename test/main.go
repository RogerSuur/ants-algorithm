package main

import (
	"fmt"
)

func main() {
	employee := map[string]int{
		"Mark": 10, "Sandy": 20,
		"Rocky": 30, "Rajiv": 40, "Kate": 50,
	}
	for key, element := range employee {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}

	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
