package switcher

import (
	"fmt"
	"passwordGen/generator"
)

func CaseSwitcher(settings map[string]*bool, len *int) {
	for {
		choice := choicer(settings)
		switch choice {
		case 1:
			*settings["digits"] = !*settings["digits"]
		case 2:
			*settings["special"] = !*settings["special"]
		case 3:
			*settings["lowercase"] = !*settings["lowercase"]
		case 4:
			*settings["uppercase"] = !*settings["uppercase"]
		case 5:
			fmt.Print("введите новую длину: ")
			var temp int
			fmt.Scan(&temp)
			if temp <= 0 {
				fmt.Println("значение было введено не коректно")
				break
			}
			*len = temp
		case 6:
			v, _ := checkSwitcher(settings)
			if v {
				fmt.Printf("ваш пароль: %s \n", generator.GenPassword(settings, *len))
			}
		case 7:
			fmt.Print("До свидания")
			return
		default:
			fmt.Println("неправильный выбор")
		}
		fmt.Println("—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–")
		showState(settings, len)
	}
}
