package shapes

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c *Circle) SetCircle(radius float64) {
	c.radius = radius
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) String() string {
	return fmt.Sprintf("текущий (r = %.2f)", c.radius)
}
