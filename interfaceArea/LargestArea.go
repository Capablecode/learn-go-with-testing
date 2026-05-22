package interfacearea

import (
	"fmt"
	"reflect"
)

type Shape interface {
	Area() int
}

type Rectangle struct {
	Width  int
	Height int
}
type Square struct {
	Side int
}

func (rr Rectangle) Area() int {
	return rr.Width * rr.Height
}

func (ss Square) Area() int {
	return ss.Side * ss.Side
}

func largestArea(shapes []Shape) int {
	maxArea := 0
	getStructName := ""
	for _, value := range shapes {
		checkMaxArea := value.Area()
		if checkMaxArea > maxArea {
			maxArea = checkMaxArea
			getStructName = reflect.TypeOf(value).Name()
		}

	}
	fmt.Println(getStructName)
	return maxArea
}

func main() {
	rectangleInstance := Rectangle{20, 10}
	SquareInstance := Square{5}
	fmt.Println(largestArea([]Shape{rectangleInstance, SquareInstance}))
}
