package main

import (
	"AreaCalculator/shapes"
	"AreaCalculator/switcher"
	"fmt"
)

func main() {
choiceFigure:
	var choiceMenu int
	fmt.Println("Выберите фигуру:")
	fmt.Println("1) Круг")
	fmt.Println("2) Прямоугольник")
	fmt.Println("3) Квадрат")
	fmt.Println("4) Треугольник")
	fmt.Println("5) Выход")
	fmt.Print("ваш выбор : ")
	fmt.Scan(&choiceMenu)
	switch choiceMenu {
	case 1:
		c := shapes.Circle{}
		switcher.ShapeSwitcher(&c)
		goto choiceFigure
	case 2:
		r := shapes.Rectangle{}
		switcher.ShapeSwitcher(&r)
		goto choiceFigure
	case 3:
		s := shapes.Square{}
		switcher.ShapeSwitcher(&s)
		goto choiceFigure
	case 4:
		t := shapes.Triangle{}
		switcher.ShapeSwitcher(&t)
		goto choiceFigure
	case 5:
		break
	default:
		fmt.Println("неправильный выбор")
	}
}
