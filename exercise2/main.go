package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(stream Stream, tweets chan Tweet, wg *sync.WaitGroup) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			wg.Done()
			return
		}
		wg.Add(1)
		tweets <- *tweet
	}
}

func consumer(tweets chan Tweet, wg *sync.WaitGroup) {
	for t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
		wg.Done()
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	wg := &sync.WaitGroup{}
	tweets := make(chan Tweet)
	wg.Add(1)
	go producer(stream, tweets, wg)
	go consumer(tweets, wg)
	wg.Wait()
	fmt.Printf("Process took %s\n", time.Since(start))
}