package main

import (
	"context"
	"findNumInArr/find"
	"findNumInArr/generation"
	"flag"
	"fmt"
	"sync"
)

func main() {
	quantityArr := flag.Int("arr", 5, "кол-во масивов")
	quantityElement := flag.Int("element", 100, "кол-во элементов в маисве")
	findNum := flag.Int("find", 50, "число которое надо найти")
	flag.Parse()
	arrays := generation.GenerationArr(*quantityArr, *quantityElement)
	var wg sync.WaitGroup
	resultChan := make(chan find.Result)
	ctx, cancel := context.WithCancel(context.Background())
	found := false
	defer cancel()
	for i := 0; i < len(arrays); i++ {
		wg.Add(1)
		go find.SearchArray(ctx, i, arrays[i], *findNum, resultChan, &wg)
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()
	for result := range resultChan {
		fmt.Printf("Найдено Массив %d, позиция %d: %d\n",
			result.ArrayIndex, result.Position, result.Value)
		found = true
		cancel()
		break
	}

	if !found {
		fmt.Println("число не найдено")
	}
	fmt.Println("gg")
}
