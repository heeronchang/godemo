package main

import (
	"log"
	"sync/atomic"
)

func main() {
	sem := make(chan struct{}, 10)
	var count int32

	for {
		go func() {
			sem <- struct{}{}
			atomic.AddInt32(&count, 1)
			defer func() {
				<-sem
				atomic.AddInt32(&count, -1)
			}()

			log.Printf("count is :%d\n", count)
		}()
	}
}
