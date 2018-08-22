package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var mCond *sync.Cond
var wait bool
var wg sync.WaitGroup

func pullResource(url string) {
	mCond.L.Lock()
	for wait {
		mCond.Wait()
	}
	mCond.L.Unlock()
	fmt.Println(time.Now(), "|Pull resource from: ", url)

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
	fmt.Println(time.Now(), "|Done with resource: ", url)
	wg.Done()
}

func main() {
	url := flag.String("url", "http://test-service/bucket/content-01.ts", "resource to pull from")
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
