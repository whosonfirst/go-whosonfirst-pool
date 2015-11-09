package pool

import (
	"sync"
	"sync/atomic"
)

// https://github.com/SimonWaldherr/golang-examples/blob/2be89f3185aded00740a45a64e3c98855193b948/advanced/lifo.go

func NewPool() *Pool {
	return &Pool{mutex: &sync.Mutex{}}
}

type Pool struct {
	nodes []interface{}
	count int64
	mutex *sync.Mutex
}

func (pl *Pool) Length() int64 {
	return pl.count
}

func (pl *Pool) Push(i interface{}) {
	pl.nodes = append(pl.nodes[:pl.count], i)
	atomic.AddInt64(&pl.count, 1)
}

func (pl *Pool) Pop() interface{} {

	if pl.count == 0 {
		return 0
	}

	pl.mutex.Lock()

	atomic.AddInt64(&pl.count, -1)
	i := pl.nodes[pl.count]

	pl.mutex.Unlock()
	return i
}
