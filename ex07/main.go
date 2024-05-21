package main

import (
	"fmt"
	"sync"
	"time"
)

var or func(channels ...<-chan interface{}) <-chan interface{}

func main() {
	or := func(channels ...<-chan interface{}) <-chan interface{} {
		out := make(chan interface{}, 1)

		wg := sync.WaitGroup{}
		wg.Add(1)

		for _, channel := range channels {
			go func(channel <-chan interface{}) {
				select {
				case <-channel:
					defer wg.Done()
					out <- struct{}{}
				}
			}(channel)
		}

		wg.Wait()
		close(out)

		return out
	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v", time.Since(start))
}
