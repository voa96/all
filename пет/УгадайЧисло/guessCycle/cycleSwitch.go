package guesscycle

import (
	"fmt"
	"guessNumber/changeSettings"
	"guessNumber/random"
)

func Switcher(settings map[string]*int) {
menuLoop:
	for {
		var choice string
		fmt.Println(`Хотите сыграть в игру? ("y" | "n")`)
		fmt.Print("ваш выбор: ")
		fmt.Scan(&choice)
		switch choice {
		case "y", "Y", "YES", "yes", "н", "Н":
			for {
				var choice1 int
				fmt.Printf(`	
–—–—–—–—–—–—–—–—–—–—–—–—–—–—
	текушие настрйки:    
	минимальное число: %d
	максимальное число: %d
	попыток: %d
–—–—–—–—–—–—–—–—–—–—–—–—–—–—
			`, *settings["min"], *settings["max"], *settings["attempts"])
				fmt.Println("выберите действие:")
				fmt.Println("1) изменить настройки")
				fmt.Println("2) начать с текущими настройками")
				fmt.Println("3) выйти в меню")
				fmt.Print("ваш выбор: ")
				fmt.Scan(&choice1)
				switch choice1 {
				case 1:
					changeSettings.SettingsSwitch(settings)
				case 2:
					*settings["findNumber"] = random.RandNum(*settings["min"], *settings["max"])
					startGame(settings)
				case 3:
					goto menuLoop
				default:
					fmt.Println("неправельный выбор")
				}
			}
		case "n", "N", "NO", "no", "т", "Т":
			return
		}

	}
}

func startGame(settings map[string]*int) {
	rangeNum := map[string]int{
		"min":     *settings["min"],
		"max":     *settings["max"],
		"findNum": *settings["findNumber"],
		"win":     0,
	}
	fmt.Printf("Привет, тебе предстоит угадать число в диапазоне от %d до %d \n", rangeNum["min"], rangeNum["max"])
	fmt.Printf("у тебя есть не больше %d попыток \n", *settings["attempts"])
	fmt.Println("твоя задача писать число в диапазоне, а я буду подсказывать число больше или меньше загадываемого")
	var check int
	for i := 1; i <= *settings["attempts"]; i++ {
		fmt.Printf("попытка %d/%d\n", i, *settings["attempts"])
		fmt.Printf("текуший диапазон от %d до %d \n", rangeNum["min"], rangeNum["max"])
		fmt.Print("ваше число: ")
		fmt.Scan(&check)
		fmt.Println()
		if check < rangeNum["min"] || check > rangeNum["max"] {
			fmt.Printf("твое число в не диапазона от %d до %d \n", rangeNum["min"], rangeNum["max"])
			i--
			continue
		}
		checker(&rangeNum, check)
		if rangeNum["win"] == 1 {
			break
		}
	}
	if rangeNum["win"] == 0 {
		fmt.Println("–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—")
		fmt.Printf("ты проиграл число было: %d\n", rangeNum["findNum"])
		fmt.Println("–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—")
	}
}

func checker(rangeNum *map[string]int, check int) {
	if check == (*rangeNum)["findNum"] {
		fmt.Println("–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—")
		fmt.Printf("!!!поздравляю ты угадал число %d!!!\n", (*rangeNum)["findNum"])
		fmt.Println("–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—–—")
		(*rangeNum)["win"] = 1
		return
	}

	if check < (*rangeNum)["findNum"] {
		fmt.Println("твое число меньше загаданного")
		if check > (*rangeNum)["min"] {
			(*rangeNum)["min"] = check
		}
	} else {
		fmt.Println("твое число больше загаданного")
		if check < (*rangeNum)["max"] {
			(*rangeNum)["max"] = check
		}
	}
}
