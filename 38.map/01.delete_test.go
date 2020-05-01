package _8_map

import (
	"log"
	"runtime"
	"testing"
)

var lastTotalFreed uint64
var intMap map[int]int
var cnt = 8192

func TestMem(t *testing.T) {
	printMemStats()

	initMap()
	runtime.GC()
	printMemStats()

	log.Println(len(intMap))
	for i := 0; i < cnt; i++ {
		delete(intMap, i)
	}
	log.Println(len(intMap))

	runtime.GC()
	printMemStats()

	intMap = nil
	runtime.GC()
	printMemStats()
}

func initMap() {
	intMap = make(map[int]int, cnt)

	for i := 0; i < cnt; i++ {
		intMap[i] = i
	}
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v TotalAlloc = %v  Just Freed = %v Sys = %v NumGC = %v\n",
		m.Alloc/1024, m.TotalAlloc/1024, ((m.TotalAlloc-m.Alloc)-lastTotalFreed)/1024, m.Sys/1024, m.NumGC)

	lastTotalFreed = m.TotalAlloc - m.Alloc
}
