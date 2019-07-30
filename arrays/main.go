package main

import "fmt"

func main() {
	//integer array
	var intArr [5]int
	fmt.Println(intArr)

	//boolean array
	var boolArr [3]bool
	fmt.Println(boolArr)

	//string array
	var strArr [10]string
	fmt.Println(strArr)

	//changing values of arrays
	intArr[0] = 5
	intArr[1] = 6
	fmt.Println(intArr[0])
	fmt.Println(intArr[1])

	a := [5]int{1,2,3,4,5}

	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(len(intArr))
	fmt.Println(len(boolArr))
	fmt.Println(len(strArr))

	//multidimensional array
	var aa [3][3]int
	fmt.Println(aa)

	//populating multidimensional array
	count:=0
	for i:=0;i<len(aa);i++{
		for j:=0;j<len(aa[i]);j++{
			aa[i][j] = count
			count += 1
		}
	}
	fmt.Println(aa)
}
