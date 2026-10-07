package main

import (
	"concurrency/miner"
	"concurrency/postman"
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var coal atomic.Int64
	var mails []string
	mtx := sync.Mutex{}
	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())
	coalTransferPoint := miner.MinerPool(minerContext, 3)
	mailTransferPoint := postman.PostmanPoll(postmanContext, 3)

	initTime := time.Now()

	go func() {
		time.Sleep(3 * time.Second)
		minerCancel()
	}()
	go func() {
		time.Sleep(6 * time.Second)
		postmanCancel()
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range coalTransferPoint {
			coal.Add(int64(v))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range mailTransferPoint {
			mtx.Lock()
			mails = append(mails, v)
			mtx.Unlock()
		}
	}()
	wg.Wait()
	fmt.Println("coal: ", coal.Load())
	mtx.Lock()
	fmt.Println("mail: ", len(mails))
	mtx.Unlock()
	fmt.Println(time.Since(initTime))
}

пн 11ч
492 