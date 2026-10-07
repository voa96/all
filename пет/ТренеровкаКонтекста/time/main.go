package main

import (
	"context"
	"contextTrain/timer"
	"flag"
	"fmt"
	"time"
)

func main() {
	timeP := flag.Int("time", 10, "Время таймера в секундах")
	parentContext, parentCancel := context.WithCancel(context.Background())
	defer parentCancel()
	go func() {
		fmt.Println("нажмите  Enter для остановки")
		fmt.Scanln()
		parentCancel()
		fmt.Println("отменено пользователем")
	}()
	time.Sleep(100 * time.Millisecond)
	timer.CountDown(parentContext, *timeP)
}
