package hpool

import "github.com/hqpko/hbuffer"

func NewBufferPool(poolSize, maxBufferSize int) *Pool {
	p := NewPool(
		func() interface{} {
			return hbuffer.NewBuffer()
		},
		poolSize,
	)

	p.SetPutChecker(func(i interface{}) bool {
		bf, ok := i.(*hbuffer.Buffer)
		if !ok || bf == nil {
			return true
		}
		bf.Reset()
		return bf.Cap() > maxBufferSize
	})

	return p
}
