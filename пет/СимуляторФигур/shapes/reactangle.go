package shapes

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) SetRectangle(width, height float64) {
	r.height = height
	r.width = width
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Квадрат  (Width : %.2f, Height : %.2f)", r.width, r.height)
}
