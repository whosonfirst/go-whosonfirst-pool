package pool

import (
	"strconv"
	"sync"
	"sync/atomic"
)

// https://github.com/SimonWaldherr/golang-examples/blob/2be89f3185aded00740a45a64e3c98855193b948/advanced/lifo.go

type PoolItem interface {
	StringValue() string
	IntValue() int64
}

type PoolInt struct {
	PoolItem
	Int int64
}

func (i PoolInt) StringValue() string {
	return strconv.FormatInt(i.Int, 10)
}

func (i PoolInt) IntValue() int64 {
	return i.Int
}

type PoolString struct {
	PoolItem
	String string
}

func (s PoolString) StringValue() string {
	return s.String
}

func (s PoolString) IntValue() int64 {
	return int64(0)
}

func NewPool() *Pool {
	return &Pool{mutex: &sync.Mutex{}}
}

type Pool struct {
	nodes []PoolItem
	count int64
	mutex *sync.Mutex
}

func (pl *Pool) Length() int64 {
	return pl.count
}

func (pl *Pool) Push(i PoolItem) {
	pl.nodes = append(pl.nodes[:pl.count], i)
	atomic.AddInt64(&pl.count, 1)
}

func (pl *Pool) Pop() (PoolItem, bool) {

	if pl.count == 0 {
		return nil, false
	}

	pl.mutex.Lock()

	atomic.AddInt64(&pl.count, -1)
	i := pl.nodes[pl.count]

	pl.mutex.Unlock()
	return i, true
}
