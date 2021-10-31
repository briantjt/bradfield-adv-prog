package idservice

import (
	"sync"
	"sync/atomic"
)

type IdService interface {
	// Returns values in ascending order; it should be safe to call
	// GetNext() concurrently without any additional synchronization.
	GetNext() uint64
}

type IdGenMutex struct {
	mu    sync.Mutex
	count uint64
}

func NewIdGenMutex() *IdGenMutex {
	return &IdGenMutex{mu: sync.Mutex{}, count: 0}
}

func (g *IdGenMutex) GetNext() uint64 {
	g.mu.Lock()
	defer g.mu.Unlock()
	v := g.count
	g.count += 1
	return v
}

type IdGenAtomic struct {
	count uint64
}

func NewIdGenAtomic() *IdGenAtomic {
	return &IdGenAtomic{}
}

func (g *IdGenAtomic) GetNext() uint64 {
	v := atomic.AddUint64(&g.count, 1)
	return v - 1
}

type IdGenChan struct {
    in chan struct{}
    out chan uint64
}

func NewIdGenChan() *IdGenChan {
    idGen := &IdGenChan{in: make(chan struct{}), out: make(chan uint64)}
    go func() {
        var v uint64 = 0
        for {
            <- idGen.in
            old := v
            v += 1
            idGen.out <- old
        }
    }()
    return idGen
}

func (g *IdGenChan) GetNext() uint64 {
    g.in <- struct{}{}
    return <- g.out
}
