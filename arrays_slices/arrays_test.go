package main

import (
	"reflect"
	"testing"
)

func Test_Sum(t *testing.T){
	

	t.Run("Sum of collection of any size", func(t *testing.T){
		numbers := []int{1, 2, 3}
		
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d, want %d, %v", got, want, numbers)
		}
	})
}

func Test_SumAll(t *testing.T){

	t.Run("return Slices of Sums of any number of Slices", func(t *testing.T) {

		sliceOfSlices := [][]int{
			{1,2,3},
			{4,5,6},
		}

		got := sumAll(sliceOfSlices)
		want := []int{6,15}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})

}

func Test_SumTails(t *testing.T){

	checkSumTailEqual := func(t testing.TB, got, want []int){
		t.Helper()
	
		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v, want %v", got, want)
		}
	}


	t.Run("return slice with all the tail values only", func(t *testing.T){

		got := sumAllTails([][]int{{1,2,3},{4,5,6}})
		want := []int{5,11}

		checkSumTailEqual(t, got, want)

	})

	t.Run("safely sum empty slices", func(t *testing.T){

		got := sumAllTails([][]int{{1,2,3},{}})
		want := []int{5,0}

		checkSumTailEqual(t, got, want)

	})

}

