package main

import "fmt"

func main() {
	array := [3]int8{2, 3, 4}
	// fmt.Printf("origin:\t\t")
	for _, v := range array {
		fmt.Printf("%d(%p) ", v, &v)
		v += 5
	}
	fmt.Println()
	for _, v := range array {
		fmt.Printf("%d(%p) ", v, &v)
	}
	fmt.Println()
	// fmt.Printf("\nfn(changed):\t")
	// arr(array)
	// fmt.Printf("\norigin:\t\t")
	// for _, v := range array {
	// fmt.Printf("%d(%p) ", v, &v)
	// }
	// fmt.Println()
	// fmt.Println()
	// fmt.Println(array)
	// arrlink(&array)
}

func arrlink(array *[3]int8) {
	fmt.Println(&array)
}

func arr(array [3]int8) {
	for _, v := range array {
		v += 5
		fmt.Printf("%d(%p) ", v, &v)
	}
}
