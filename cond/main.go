package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"time"
)

var status int64

func main() {

	cond := &sync.Cond{L: &sync.Mutex{}}
	for i := 0; i < 10; i++ {
		go listen(cond)
	}
	time.Sleep(1 * time.Second)
	go broadcast(cond)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

func broadcast(cond *sync.Cond) {
	cond.L.Lock()
	defer cond.L.Unlock()
	atomic.StoreInt64(&status, 1)
	cond.Broadcast()
}

func listen(cond *sync.Cond) {
	cond.L.Lock()
	defer cond.L.Unlock()
	for atomic.LoadInt64(&status) != 1 {
		cond.Wait()
	}
	log.Println("listen ok")
}
