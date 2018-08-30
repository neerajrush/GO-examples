package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
	"math/rand"
	"io"
)

var mCond *sync.Cond
var wait bool
var wg sync.WaitGroup

var delayThreads int
var firstThread bool
func pullResource(url string) {
	mCond.L.Lock()
	for wait {
		mCond.Wait()
	}
	sleepFor := delayThreads
	if firstThread {
		sleepFor = 0
		firstThread = false
	}
	mCond.L.Unlock()
	sleepFor += rand.Inin(10)
	time.Sleep(time.Duration(sleepFor * int(time.Millisecond)))
	
	fmt.Println(time.Now(), "|Pull resource from: ", url, " after sleeping ", sleepFor, "ms")

	var c = &http.Client{}

	rsp, err := c.Get(url)
	if err == nil && rsp.StatusCode == http.StatusOK {
		fmt.Println(time.Now(), "|ResponseCode:", rsp.Status)
	} else {
		if rsp != nil {
			fmt.Println(time.Now(), "|ERROR: Resp: ", rsp.Status)
		} else {
			fmt.Println(time.Now(), "|ERROR: Resp(err): ", string(err.Error()))
		}
	}
	dBuf := make([]byte, 8192*8192)
	var i, total int
	for err != io.EOF && err == nil && rsp.Body != nil {
		i, err = rsp.Body.Read(dBuf)
		total += i
	}
	fmt.Println(time.Now(), "|Read total bytes: ", total)
	fmt.Println(time.Now(), "|Done with resource: ", url)
	wg.Done()
}

func main() {
	url := flag.String("url", "http://test-service/bucket/content-01.ts", "resource to pull from")
	delayThreads = *flag.Int("delay", 20, "delay second and other threads in ms")
	flag.Parse()
	mCond = sync.NewCond(&sync.Mutex{})
	wait = true

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go pullResource(*url)
	}
	time.Sleep(3 * time.Second)
	wait = false
	mCond.Broadcast()

	wg.Wait()
}
