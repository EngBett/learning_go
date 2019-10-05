package main

import "fmt"

func main() {
	// similar to range in python
	strArr := []string{"a", "b", "c", "d"}

	for i, c := range strArr {
		fmt.Println(i, c)
	}

	fmt.Println("")

	// can also range over key value pairs in maps
	m := map[string]int{"a": 12, "b": 28, "c": 45}

	for key, val := range m {
		fmt.Println("Key:", key, "Val:", val)
	}

	fmt.Println("")

	//can also range over keys in maps
	for k := range m {
		fmt.Println(k)
	}

	fmt.Println("")

	//values only
	for k := range m {
		fmt.Println(m[k])
	}

}
