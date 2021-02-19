package scanstring

import (
	"runtime"
	"sync"
)

func chunkSlice(arr []byte) (chunks [][]byte, size int) {
	size = (len(arr) + runtime.NumCPU() - 1) / runtime.NumCPU()
	chunks = make([][]byte, len(arr)/size)

	for i := range chunks {
		chunks[i] = arr[size*i : size*i+size]
	}

	return
}

type CounterChannel = chan struct{}

func scanChunkIntoChannel(chunk []byte, wantedByte byte, ch CounterChannel, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, b := range chunk {
		if b == wantedByte {
			ch <- struct{}{}
		}
	}
}

func countByteOccurencesChannel(src []byte, wantedByte byte) uint {
	chunks, _ := chunkSlice(src)

	var wg sync.WaitGroup

	cch := make(CounterChannel, len(src))

	for _, chunk := range chunks {
		wg.Add(1)
		go scanChunkIntoChannel(chunk, wantedByte, cch, &wg)
	}
	wg.Wait()

	return uint(len(cch))
}

type SafeCounter struct {
	mu sync.Mutex
	c  uint
}

func scanChunkToSafeCounter(chunk []byte, wantedByte byte, sc *SafeCounter, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, b := range chunk {
		if b == wantedByte {
			sc.mu.Lock()
			sc.c += 1
			sc.mu.Unlock()
		}
	}
}

func countByteOccurencesSyncConcurrency(src []byte, wantedByte byte) uint {
	chunks, _ := chunkSlice(src)
	sc := SafeCounter{
		c: 0,
	}

	var wg sync.WaitGroup

	for _, chunk := range chunks {
		wg.Add(1)
		go scanChunkToSafeCounter(chunk, wantedByte, &sc, &wg)
	}
	wg.Wait()

	return sc.c
}

func countByteOccurencesSync(src []byte, wantedByte byte) uint {
	count := uint(0)
	for _, b := range src {
		if b == wantedByte {
			count++
		}
	}

	return count
}
