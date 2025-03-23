package main

import (
	"fmt"
)

func Sum(numbers []int) int {

	sum := 0

	for _, i := range(numbers) {
		sum += i;
	}

	return sum

}

func sumAll(sliceOfSlices [][]int) []int {

	var finalSliceOfInts []int;

	for _, sliceOfInts := range(sliceOfSlices){
		finalSliceOfInts = append(finalSliceOfInts, Sum(sliceOfInts))
	}

	return finalSliceOfInts

}

func sumAllTails(sliceOfSlices [][]int) []int {

	var finalSlicesOfTails []int;

	for _, sliceOfInts := range(sliceOfSlices){

		if len(sliceOfInts) == 0 {
			finalSlicesOfTails = append(finalSlicesOfTails,0);	
			continue
		}
		finalSlicesOfTails = append(finalSlicesOfTails,Sum(sliceOfInts[1:]));
	}

	return finalSlicesOfTails
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	slice := numbers[1:]

    // Print the address of the first element in the slice
    fmt.Printf("Address of the first element: %p\n", &numbers[1])

	fmt.Printf("Address of the first element: %p\n", &slice[0])

	    // Print the memory reference of the slice descriptor itself
		fmt.Printf("Address of the slice descriptor: %p\n", slice)

    // Print the memory reference of the slice descriptor itself
    fmt.Printf("Address of the slice descriptor: %p\n", numbers)
}