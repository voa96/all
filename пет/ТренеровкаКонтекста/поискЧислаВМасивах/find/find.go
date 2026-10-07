package find

import (
	"context"
	"fmt"
	"sync"
)

type Result struct {
	ArrayIndex int
	Position   int
	Value      int
}

func SearchArray(ctx context.Context, index int, array []int, target int, resultChan chan Result, wg *sync.WaitGroup) {
	fmt.Printf("поиск в масиве %d начался\n", index)
	defer wg.Done()
	for i, val := range array {
		select {
		case <-ctx.Done():
			return
		default:
			if val == target {
				select {
				case resultChan <- Result{
					ArrayIndex: index,
					Position:   i,
					Value:      val,
				}:
				case <- ctx.done()
					fmt.Printf("массив %d отправил результат\n", index)
				default:
					fmt.Printf("массив %d: канал занят, результат не отправлен\n", index)
				}
				return
			}
		}
	}
}
