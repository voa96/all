package shapes

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

func (s *Square) SetSquare(side float64) {
	s.side = side
}

func (s Square) Area() float64 {
	return math.Pow(s.side, 2)
}

func (s Square) Perimeter() float64 {
	return 4 * s.side
}

func (s Square) String() string {
	return fmt.Sprintf("Квадрат (side : %.2f)", s.side)
}
