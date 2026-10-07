package timer

import (
	"context"
	"fmt"
	"time"
)

func CountDown(ctx context.Context, timeP int) {
	for i := timeP; i >= 1; i-- {
		select {
		case <-ctx.Done():
			fmt.Printf("Стоп, причина : %v\n", ctx.Err())
			return
		default:
			fmt.Printf("%d ", i)
			time.Sleep(1 * time.Second)
		}
	}
}
