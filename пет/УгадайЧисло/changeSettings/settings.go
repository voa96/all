package changeSettings

import (
	"fmt"
)

func SettingsSwitch(settings map[string]*int) {
	for {
		var choice int
		fmt.Printf(`
–—–—–—–—–—–—–—–—–—–—–—–—–—–—
	Выберите что хотите изменить: 
	1)минимальное число (текущее %d)
	2)максимальное число (текущее %d)
	3)число попыток (текущее %d)
	4)Выйти
–—–—–—–—–—–—–—–—–—–—–—–—–—–—
		`, *settings["min"], *settings["max"], *settings["attempts"])
		fmt.Print("ваш выбор: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var (
				newMin int
			)
			newMax := *settings["max"]
			fmt.Print("введите новое минимальное число: ")
			fmt.Scan(&newMin)
			newMin, newMax = validateRange(newMin, *settings["max"])
			*settings["min"] = newMin
			*settings["max"] = newMax
		case 2:
			var (
				newMin int
			)
			newMax := *settings["max"]
			fmt.Print("введите новое максимальное число: ")
			fmt.Scan(&newMax)
			newMin, newMax = validateRange(*settings["min"], newMax)
			*settings["min"] = newMin
			*settings["max"] = newMax
		case 3:
			var newAttempts int
			fmt.Print("введите новое количество попыток: ")
			fmt.Scan(&newAttempts)
			if newAttempts <= 0 {
				fmt.Println("кол-во попыток не должно быть меньше 0")
				break
			}
			*settings["attempts"] = newAttempts
		case 4:
			return
		default:
			fmt.Println("неправельный выбор")
		}
	}
}

func validateRange(min, max int) (int, int) {

	if max < min {
		max, min = min, max
		fmt.Println("максимальное число не должно быть меньше минимального")
		fmt.Printf("новое минимальное число: %d\n", min)
		fmt.Printf("новое максимальное число: %d\n", max)
	}
	return min, max
}
