package main

import (
	"context"
	"log"
	"os"
	"time"
)

func doStuff2(ctx context.Context) {
	for {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			log.Println("done")
			return
		default:
			log.Println("work")
		}
	}
}

func doStuff1(ctx context.Context) {
	go doStuff2(ctx)
}

func main() {
	log.SetOutput(os.Stdout)
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff1(ctx)

	time.Sleep(10 * time.Second)
	cancel()
	for {
		time.Sleep(time.Second)
	}
}
