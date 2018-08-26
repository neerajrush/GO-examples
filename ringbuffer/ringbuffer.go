package ringbuffer

import (
	"fmt"
	"sync"
	"time"
)

type RingBuffer struct {
	Mutex       sync.RWMutex
	RBuff       []string
	Size        int
	ReaderIndex int
	WriterIndex int
}

const (
	RING_BUFFER_SIZE = 1024 * 1024
)

func NewRingBuffer() *RingBuffer {
	ringBuffer := RingBuffer{
		Mutex:       sync.RWMutex{},
		RBuff:       make([]string, RING_BUFFER_SIZE, RING_BUFFER_SIZE),
		Size:        RING_BUFFER_SIZE,
		ReaderIndex: 0,
		WriterIndex: 0,
	}

	return &ringBuffer
}

func (rb *RingBuffer) WriteData(msg string) bool {
	rb.Mutex.Lock()
	defer rb.Mutex.Unlock()

	if rb.WriterIndex+1 == rb.ReaderIndex {
		return false
	}

	if rb.WriterIndex+1 == RING_BUFFER_SIZE {
		if rb.ReaderIndex == 0 {
			return false
		}
	}

	rb.RBuff[rb.WriterIndex] = msg

	if rb.WriterIndex+1 == RING_BUFFER_SIZE {
		rb.WriterIndex = 0
	} else {
		rb.WriterIndex++
	}

	return true
}

func (rb *RingBuffer) ReadData() (string, bool) {
	rb.Mutex.RLock()
	defer rb.Mutex.RUnlock()

	if rb.ReaderIndex == rb.WriterIndex {
		return "", false
	}

	result := rb.RBuff[rb.ReaderIndex]

	rb.ReaderIndex++

	if rb.ReaderIndex == RING_BUFFER_SIZE {
		rb.ReaderIndex = 0
	}

	return result, true
}

func (rb *RingBuffer) Write(id string, count int, wg *sync.WaitGroup) {
	wIdx := 1
	for wIdx < count {
		msg := fmt.Sprintf("%s(%d)", id, wIdx)
		ok := rb.WriteData(msg)
		if ok {
			wIdx++
		}
		time.Sleep(100 * time.Millisecond)
	}
	wg.Done()
}

func (rb *RingBuffer) Read(id string, count int, wg *sync.WaitGroup) {
	rIdx := 1
	for rIdx < count {
		if msg, ok := rb.ReadData(); ok {
			fmt.Println(id, " ==> ", msg)
			rIdx++
		}
		time.Sleep(100 * time.Millisecond)
	}
	wg.Done()
}
