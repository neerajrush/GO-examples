Created ringbuffer as package so we can importit in other applications. The example is provided below.


Here is the test results of ringbuffer package.

* #go build ringbuffer
* #go install ringbuffer
* #go test -v ringbuffer

```
go test -v ringbuffer
=== RUN   TestSetup
--- PASS: TestSetup (0.00s)
=== RUN   TestTeardown
--- PASS: TestTeardown (0.00s)
=== RUN   TestWrite
--- PASS: TestWrite (0.00s)
=== RUN   TestRead
--- PASS: TestRead (0.00s)
PASS
Reader2  ==>  Writer2(1)
Reader1  ==>  Writer1(1)
ok  	ringbuffer	(cached)
```

Here is the example code that imports it as package.
my-app/cmd/main.go:
```
package main

import (
	"fmt"
	"ringbuffer"
	"sync"
)

func main() {
	ringBuff := ringbuffer.NewRingBuffer()
	wg := &sync.WaitGroup{}
	if ok := ringBuff.WriteData("test1"); ok {
		fmt.Println("Data written: test1")
	}

	if msg, ok := ringBuff.ReadData(); ok {
		fmt.Println("Data Read:", msg)
	}

	go ringBuff.Write("Writer1", 100, wg)
	wg.Add(1)
	go ringBuff.Write("Writer2", 100, wg)
	wg.Add(1)

	go ringBuff.Read("Reader1", 100, wg)
	wg.Add(1)
	go ringBuff.Read("Reader2", 100, wg)
	wg.Add(1)

	wg.Wait()
}
```

#go build
#./cmd

```
Data written: test1
Data Read: test1
Reader2  ==>  Writer1(1)
Reader1  ==>  Writer1(1)
Reader1  ==>  Writer1(2)
Reader2  ==>  Writer2(1)
Reader1  ==>  Writer2(2)
...
...
...
Reader2  ==>  Writer2(94)
Reader1  ==>  Writer1(95)
Reader2  ==>  Writer2(96)
Reader2  ==>  Writer1(96)
Reader1  ==>  Writer1(96)
```


