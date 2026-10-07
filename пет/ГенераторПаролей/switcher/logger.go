package switcher

import (
	"errors"
	"fmt"
)

func checkSwitcher(settings map[string]*bool) (bool, error) {
	quantity := 0
	for _, v := range settings {
		if *v {
			quantity++
		}
	}
	if quantity > 0 {
		return true, nil
	}
	return false, errors.New("!!!нужно выбрать хотябы один пункт!!!")
}
func showState(settings map[string]*bool, len *int) {
	fmt.Printf("длина пароля : %d \n", *len)
	fmt.Println("текуший пароль будет состоять из : ")
	if *settings["digits"] {
		fmt.Println("цифры")
	}
	if *settings["special"] {
		fmt.Println("специальные символы")
	}
	if *settings["lowercase"] {
		fmt.Println("букв в нижнем регистре")
	}
	if *settings["uppercase"] {
		fmt.Println("букв в верхнем регистре")
	}
	_, err := checkSwitcher(settings)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println("—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–")
}

func choicer(settings map[string]*bool) int {
	var choice int
	action := map[bool]string{
		true:  "удалить из пароля",
		false: "добавить в пароль",
	}
	fmt.Println("выберите число для действия:")
	fmt.Printf("1) %s цифры \n", action[*settings["digits"]])
	fmt.Printf("2) %s специальные символы \n", action[*settings["special"]])
	fmt.Printf("3) %s буквы нижнего регистра\n", action[*settings["lowercase"]])
	fmt.Printf("4) %s буквы вернего регистра\n", action[*settings["uppercase"]])
	fmt.Printf("5) изменить длину \n")
	fmt.Printf("6) сгенерировать пароль\n")
	fmt.Printf("7) Выйти\n")
	fmt.Println("—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–")
	fmt.Print("ваш выбор: ")
	fmt.Scan(&choice)
	fmt.Println("—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–")
	return choice
}
