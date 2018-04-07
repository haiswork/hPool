package hpool

import (
	"sync/atomic"
)

type Pool struct {
	newFunc  func() interface{}
	poolChan chan interface{}

	getCount    uint32
	putCount    uint32
	newCount    uint32
	removeCount uint32
	debug       bool
}

func NewPool(newFunc func() interface{}, poolSize int) *Pool {
	return &Pool{
		newFunc:  newFunc,
		poolChan: make(chan interface{}, poolSize),
	}
}

func (p *Pool) IsDebug() bool {
	return p.debug
}

func (p *Pool) SetDebug(b bool) {
	p.debug = b
}

func (p *Pool) Get() interface{} {
	if p.debug {
		atomic.AddUint32(&p.getCount, 1)
	}
	select {
	case i := <-p.poolChan:
		return i
	default:
		if p.debug {
			atomic.AddUint32(&p.newCount, 1)
		}
		return p.newFunc()
	}
}

func (p *Pool) Put(i interface{}) {
	if i == nil {
		return
	}
	if p.debug {
		atomic.AddUint32(&p.putCount, 1)
	}
	select {
	case p.poolChan <- i:
	default:
		if p.debug {
			atomic.AddUint32(&p.removeCount, 1)
		}
	}
}

func (p *Pool) GetDebugInfo() (uint32, uint32, uint32, uint32) {
	return p.newCount, p.getCount, p.putCount, p.removeCount
}
