package main

import "fmt"

func main() {

	//Maps are similar to hastables in java or dictionaries in python
	m := make(map[string]int)

	m["a"] = 12
	m["b"] = 15

	fmt.Println(m)

	//print specific value
	fmt.Println(m["a"])

	//length of map
	fmt.Println(len(m))

	//remove key pair from map
	delete(m, "a")
	fmt.Println(m)

	//coincise definition
	m2 := map[string]int{"key1": 1, "key2": 2}

	fmt.Println(m2)

	//the value and whether the value is present.
	val, is_val_present := m["b"]
	fmt.Println(val)
	fmt.Println(is_val_present)

	_, is_val_present2 := m["a"]

	fmt.Println(is_val_present2)

}
