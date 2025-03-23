package structs

import (
	"math"
	"testing"
)

func Test_Perimeter(t *testing.T){

	checkPerimeter := func(t testing.TB, shape Shape, want float64){
		t.Helper()

		got := shape.getPerimeter()
		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	}

	perimeterTests := []struct {
		name string
		shape Shape
		want float64
	}{
		{
			name: "Test Rectangle Perimeter",
			shape: Rectangle{4.0, 5.0}, 
			want:  18.00,
		},
		{
			name: "Test Circle Perimeter",
			shape: Circle{10.0},
			want:  math.Pi * 2 * 10.0,
		},
	}

	for _,tt := range(perimeterTests){

		t.Run(tt.name, func(t *testing.T) {
			checkPerimeter(t, tt.shape, tt.want)
		})

	}
}

func Test_Area(t *testing.T){

	checkArea := func(t testing.TB, shape Shape, want float64){
		t.Helper()

		got := shape.getArea()
		if got != want {
			t.Errorf("%#v got %g want %g", shape, got, want)
		}
	}

	areaTest := []struct{
		name string
		shape Shape
		want float64
	}{
		{
			name: "Test_Rectangle_Area",
			shape: Rectangle{4.0, 5.0}, 
			want:  20.00,
		},
		{
			name: "Test_Circle_Area",
			shape: Circle{10.0},
			want:  math.Pi * 10.0 * 10.0,
		},
		{
			name: "Test_Triangle_Area",
			shape: Triangle{2.0, 3.0},
			want:  3.0,
		},
	}

	for _, tt := range(areaTest){

		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.want)
		})
	}

}