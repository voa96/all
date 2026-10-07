package main

import (
	"flag"
	"passwordGen/switcher"
)

func main() {
	length := flag.Int("length", 8, "кол-во символов")
	lowercase := flag.Bool("lowercase", true, "нижний регистр")
	uppercase := flag.Bool("uppercase", true, "верхний регистр")
	digits := flag.Bool("digit", false, "цифры")
	special := flag.Bool("special", false, "символы")

	flag.Parse()

	settings := map[string]*bool{
		"digits":    digits,
		"special":   special,
		"lowercase": lowercase,
		"uppercase": uppercase,
	}
	switcher.CaseSwitcher(settings, length)
}
