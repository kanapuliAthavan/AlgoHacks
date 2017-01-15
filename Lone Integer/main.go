package main

import (
	"fmt"
	"sort"
)

func main() {

	var test int
	fmt.Scan(&test)
	loop(test)

}

func loop(times int) {
	count := 0
	mySlice := make([]int, 0, 1)
	newSlice := make([]int, 0, 1)
	for i := 0; i < times; i++ {
		var arrayLength int
		fmt.Scan(&arrayLength)
		array := make([]int, arrayLength, arrayLength+1)
		for z := 0; z < arrayLength; z++ {
			fmt.Scan(&array[z])
		}
		//Sort the array in increasing order
		sort.Ints(array)
		fmt.Println("sorted Array: ", array)
		mySlice = append(mySlice, array...)
		fmt.Println("My Slice : ", mySlice)
		for i := 0; i < len(mySlice); i++ {
			temp := mySlice[i]
			for a, b := range mySlice {
				if i != a && a >= i && temp == b {
					count++
					if count == 1 {
						count = 0
						newSlice = append(mySlice[:i+1], mySlice[i+2:]...)
						copy(mySlice, newSlice)
						fmt.Println("Deleted Slice : ", mySlice)
					}
					//newSlice = append(mySlice, temp)
					fmt.Println(count)

					//fmt.Println(mySlice)
				}
			}
		}

	}
}