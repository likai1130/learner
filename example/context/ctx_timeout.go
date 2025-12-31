package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	log.Println("start")
	signal := make(chan struct{}, 1)
	/*ctx,cancel := context.WithTimeout(context.Background(),1*time.Second)
	go exec(ctx,cancel,signal)
	<- signal*/

	go testChan(signal)
	<-signal
	log.Println("end")
}

func exec(ctx context.Context, cancel context.CancelFunc, signal chan struct{}) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("time out. err=%s\n", ctx.Err())
			return
		}
	}
	time.Sleep(1 * time.Second)
	fmt.Println("睡过")
	signal <- struct{}{}
	return

}

func testChan(signal chan struct{}) {
	fmt.Println("aaa")
	close(signal)
	context.TODO()
}
