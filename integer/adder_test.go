package integer

import (
	"fmt"
	"testing"
)

func Test_Adder(t *testing.T){

	sum := Add(2,3)
	expected := 5 

	if sum != expected {
		t.Errorf("got '%d', want '%d'", sum, expected)
	}

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}