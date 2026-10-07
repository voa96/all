package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func miner(
	ctx context.Context,
	transferPoint chan<- int,
	num int, power int,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("ШАХТ %d завершил добычу\n", num)
			return
		case <-time.After(1 * time.Second):
			fmt.Printf("я шахтер %d, я добыл %d\n", num, power)
		}
		select {
		case <-ctx.Done():
			fmt.Printf("ШАХТ %d завершил добычу\n", num)
			return
		case transferPoint <- power:
			fmt.Printf("я шахтер %d, я передал %d\n", num, power)
		}
	}
}

func MinerPool(ctx context.Context, minerCount int) <-chan int {
	coalTransfer := make(chan int)
	wg := &sync.WaitGroup{}
	for i := 1; i < minerCount; i++ {
		wg.Add(1)
		go miner(ctx, coalTransfer, i, i*10, wg)
	}
	go func() {
		wg.Wait()
		close(coalTransfer)
	}()

	return coalTransfer
}
