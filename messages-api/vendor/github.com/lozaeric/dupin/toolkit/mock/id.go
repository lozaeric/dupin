package mock

import (
	"fmt"
	"sync"
)

const idLength = 20

var (
	counter = 1
	mtx     = sync.Mutex{}
)

func GenerateValidID() string {
	mtx.Lock()
	defer mtx.Unlock()
	ID := fmt.Sprintf("%020d", counter)
	counter++
	return ID
}
