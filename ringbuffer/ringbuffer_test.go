package ringbuffer

import (
	"sync"
	"testing"
)

var ringBuff *RingBuffer
var wg *sync.WaitGroup

func TestSetup(t *testing.T) {
	ringBuff = NewRingBuffer()
	wg = &sync.WaitGroup{}
}

func TestTeardown(t *testing.T) {
	wg.Wait()
}

func TestWrite(t *testing.T) {
	go ringBuff.Write("Writer1", 100, wg)
	wg.Add(1)
	go ringBuff.Write("Writer2", 100, wg)
	wg.Add(1)
}

func TestRead(t *testing.T) {
	go ringBuff.Read("Reader1", 100, wg)
	wg.Add(1)
	go ringBuff.Read("Reader2", 100, wg)
	wg.Add(1)
}
