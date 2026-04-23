package shape

import (
	"testing"
)

// func TestPerimeter(t *testing.T) {
// 	t.Run("Perimeter of a rectangle", func(t *testing.T) {

// 		rectangle := Rectangle{10.0, 10.0}
// 		got := Perimeter(rectangle)
// 		expected := 40.0

// 		if got != expected {
// 			t.Errorf("got %.2f but expected %.2f", got, expected)
// 		}
// 	})
// }

// func TestArea(t *testing.T) {
// 	t.Run("Area of Rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{12.0, 6.0}
// 		got := rectangle.Area()
// 		expected := 72.0

// 		if got != expected {
// 			t.Errorf("got %.2f but expected %.2f", got, expected)
// 		}
// 	})

// 	t.Run("Area of a Circle", func(t *testing.T) {
// 		circle := Circle{10}
// 		got := circle.Area()
// 		expected := 314.1592653589793

// 		if got != expected {
// 			t.Errorf("got %g but expected %g", got, expected)
// 		}
// 	})

// }

// func TestArea(t *testing.T) {
// 	checkArea := func(t testing.TB, shape Shape, expected float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != expected {
// 			t.Errorf("got %g but expected %g", got, expected)
// 		}
// 	}

// 	t.Run("Rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{12.0, 6.0}
// 		checkArea(t, rectangle, 72.0)
// 	})

// 	t.Run("Cicle", func(t *testing.T) {
// 		circle := Circle{10}
// 		checkArea(t, circle, 314.1592653589793)
// 	})
// }

func TestArea(t *testing.T) {
	areaTest := []struct {
		shape   Shape
		hasArea float64
		name    string
	}{
		{name: "Rectange", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	//loop through the areaTest
	// for _, tt := range areaTest{
	// 	got := tt.shape.Area()
	// 	if got != tt.want{
	// 		t.Errorf("got %g but want %g", got, tt.want)
	// 	}
	// }

	//creating a t-run
	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g but want %g", tt.shape, got, tt.hasArea)
			}
		})
	}

}
