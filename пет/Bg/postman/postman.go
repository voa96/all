package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func postman(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- string,
	num int,
	mail string,
) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("я почтальон номер: %d, я завершил свой рабочий день \n", num)
			return
		default:
			fmt.Printf("я почтальон номер: %d начал \n", num)
			time.Sleep(1 * time.Second)
			fmt.Printf("я почтальон номер: %d, я донес письмо %s\n", num, mail)

			transferPoint <- mail
			fmt.Printf("я почтальон: %d, я завершил свой день епта\n", num)
		}
	}
}

func PostmanPoll(ctx context.Context, postmanCount int) <-chan string {
	mailTransferPoint := make(chan string)
	wg := &sync.WaitGroup{}
	for i := 1; i <= postmanCount; i++ {
		wg.Add(1)
		go postman(ctx, wg, mailTransferPoint, i, postmanToMail(i))
	}
	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()
	return mailTransferPoint
}

func postmanToMail(num int) string {
	ptm := map[int]string{
		1: "абу даби",
		2: "дабу аби",
		3: "еще что-то",
	}
	mail, ok := ptm[num]
	if !ok {
		return "лутка"
	}
	return mail
}
