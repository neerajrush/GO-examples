package  ringbuffer

import (
	"testing"
	"sync"
)

var ringBuff *RingBuffer
var wg *sync.WaitGroup

func TestSetup(t *testing.T) {
	ringBuff = NewRingBuffer()
	wg = &sync.WaitGroup{}
}

func TestWrite(t *testing.T) {
	go ringBuff.Write("Writer1", 2, wg)
	wg.Add(1)
	go ringBuff.Write("Writer2", 2, wg)
	wg.Add(1)
}

func TestRead(t *testing.T) {
	go ringBuff.Read("Reader1", 2, wg)
	wg.Add(1)
	go ringBuff.Read("Reader2", 2, wg)
	wg.Add(1)
}

func TestTeardown(t *testing.T) {
	wg.Wait()
}
