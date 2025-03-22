package iteration

import (
	"fmt"
	"testing"
)

func Test_Iteration(t *testing.T){

	t.Run("testCorrectOutput", func(t *testing.T){
		got := Repeat("a", 5)
		want := "aaaaa"
	
		if got != want {
			t.Errorf("got %q, want %q", got , want)
		}
	})


	t.Run("testCorrectLength", func(t *testing.T){
		got := len(Repeat("a", 10))
		want := 10
	
		if got != want {
			t.Errorf("got %d, want %d", got , want)
		}
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat(){
	
	repeatedString := Repeat("hello",2);
	fmt.Println(repeatedString)
	// Output: hellohello

}