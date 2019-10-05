package main

import "fmt"

func main() {
	s := make([]int, 3)
	s[1] = 3
	s[0] = 5
	s[2] = 8

	//appending
	s = append(s, 15, 18, 67)

	//length of slice
	fmt.Println(len(s))

	//part of slice
	fmt.Println(s[:3])

	fmt.Println(s[3:])

	test := make([]int, 3)

	fmt.Println(test)

	// concise slice definition
	t := []int{100, 200, 300}

	fmt.Println(t)

	//changing values

	/*fmt.Println(s)
	x:=s
	x[0]=200
	fmt.Println(x)
	fmt.Println(s)*/

	//use copy to prevent from changing both x and s
	x := make([]int, len(s))
	copy(x, s)
	x[0] = 500
	fmt.Println(x)
	fmt.Println(s)

	//2d slices
	ss := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		ss[i] = make([]int, innerLen)
		for j := 0; j < len(ss[i]); j++ {
			ss[i][j] = i + j
		}
	}

	fmt.Println(ss)

}
