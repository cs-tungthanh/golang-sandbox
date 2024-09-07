package utils

import (
	"fmt"
	"runtime"
	"time"
)

// PrintUsage print the memory usage:  go utils.PrintUsage() at main.go
func PrintUsage() {
	for {
		time.Sleep(500 * time.Millisecond)
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		// For info on each, see: https://golang.org/pkg/runtime/#MemStats
		fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
		fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
		fmt.Printf("\tHeap_inuse = %v MiB", bToMb(m.HeapInuse))
		fmt.Printf("\tFrees = %v MiB", bToMb(m.Frees))
		fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
		fmt.Printf("\tNumGC = %v\n", m.NumGC)
	}
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
