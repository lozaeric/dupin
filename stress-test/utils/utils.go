package utils

import (
	"runtime"
	"sync"
)

func parallelProcess(tasks int, worker func(wg *sync.WaitGroup)) {
	var wg sync.WaitGroup

	wg.Add(tasks)
	for i := 0; i < runtime.GOMAXPROCS(0); i++ {
		go worker(&wg)
	}
	wg.Wait()
}
