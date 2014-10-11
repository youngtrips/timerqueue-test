package main

import (
	"github.com/youngtrips/timerqueue"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int64, 1024)
	tq := timerqueue.New()
	tq.Run()
	t := make(map[int64]int64)
	//for i := 0; i < 100000; i++ {
	//tid := tq.Schedule(rand.Int63() % 60000)
	tid := tq.Schedule(2000, ch)
	t[tid] = time.Now().UnixNano()
	//}
	for {
		select {
		case id, ok := <-ch:
			if ok {
				curr := time.Now().UnixNano()
				log.Printf("time event: tid=%d, diff=%d", id, (curr-t[id])/1e6)
				//tid := tq.Schedule(rand.Int63() % 60000)
				tid := tq.Schedule(2000, ch)
				t[tid] = time.Now().UnixNano()
			}
		}
	}
}
