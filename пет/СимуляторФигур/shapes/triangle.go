package shapes

import (
	"fmt"
	"math"
)

type Triangle struct {
	a, b, c float64
}

func (t *Triangle) SetTriangle(a, b, c float64) {
	t.a = a
	t.b = b
	t.c = c
}

func (t Triangle) Area() float64 {
	p := t.Perimeter()
	return math.Sqrt((p * (p - t.a) * (p - t.b) * (p - t.c)))
}

func (t Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) String() string {
	return fmt.Sprintf("Треугольник (a: %.2f, b: %.2f, c: %.2f)", t.a, t.b, t.c)
}
