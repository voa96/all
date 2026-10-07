package switcher

import (
	"AreaCalculator/shapes"
	"fmt"
)

func ShapeSwitcher(s shapes.Shape) {
	for {
		var choice int
		fmt.Println("—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–")
		fmt.Println(s.String())
		fmt.Println("—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–")
		fmt.Println("выберите действие: ")
		fmt.Println("1) задать значения фигуре")
		fmt.Println("2) получить площадь")
		fmt.Println("3) получить Периметр")
		fmt.Println("4) Вернуться к выбору фигур")
		fmt.Print("ваш выбор: ")
		fmt.Scan(&choice)
		fmt.Println("—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–")
		switch choice {
		case 1:
			editShape(s)
		case 2:
			fmt.Println("площадь : ", s.Area())
		case 3:
			fmt.Println("периметр : ", s.Perimeter())
		case 4:
			return
		default:
			fmt.Println("неправильный выбор")
		}
	}
}

func editShape(s shapes.Shape) {
	switch v := s.(type) {
	case *shapes.Circle:
		var r float64
		fmt.Print("Введите радиус круга: ")
		fmt.Scan(&r)
		v.SetCircle(r)

	case *shapes.Rectangle:
		var w, h float64
		fmt.Print("Введите ширину: ")
		fmt.Scan(&w)
		fmt.Print("Введите высоту: ")
		fmt.Scan(&h)
		v.SetRectangle(w, h)

	case *shapes.Square:
		var side float64
		fmt.Print("Введите сторону квадрата: ")
		fmt.Scan(&side)
		v.SetSquare(side)

	case *shapes.Triangle:
		var a, b, c float64
		fmt.Print("Введите сторону A: ")
		fmt.Scan(&a)
		fmt.Print("Введите сторону B: ")
		fmt.Scan(&b)
		fmt.Print("Введите сторону C: ")
		fmt.Scan(&c)
		v.SetTriangle(a, b, c)

	default:
		fmt.Println("Неизвестная фигура!")
	}
}
