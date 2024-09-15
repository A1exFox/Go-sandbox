package main

import "fmt"

func main() {
	array := [3]int8{1, 2, 3}
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d(%p) ", array[i], &array[i])
	} // 1(0xc00009a00b) 2(0xc00009a00c) 3(0xc00009a00d)
	fmt.Println()
	change(&array)
	fmt.Printf("\n%v\n", array) // [5 6 7]
}
func change(array *[3]int8) {
	for i := 0; i < len(array); i++ {
		array[i] += 4
		fmt.Printf("%d(%p) ", array[i], &array[i])
	} // 5(0xc00009a00b) 6(0xc00009a00c) 7(0xc00009a00d)
}
