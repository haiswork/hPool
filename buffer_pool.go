package hpool

import "github.com/hqpko/hbuffer"

type BufferPool struct {
	pool *Pool
}

func NewBufferPool(poolSize, maxBufferSize int) *BufferPool {
	bufferPool := &BufferPool{
		pool: NewPool(
			func() interface{} {
				return hbuffer.NewBuffer()
			},
			poolSize,
		),
	}
	bufferPool.pool.SetPutChecker(func(i interface{}) bool {
		bf, ok := i.(*hbuffer.Buffer)
		if !ok || bf == nil {
			return true
		}
		bf.Reset()
		return bf.Cap() > maxBufferSize
	})

	return bufferPool
}

func (bf *BufferPool) Get() *hbuffer.Buffer {
	return bf.pool.Get().(*hbuffer.Buffer)
}

func (bf *BufferPool) Put(buffer *hbuffer.Buffer) {
	if buffer == nil {
		return
	}
	bf.pool.Put(buffer)
}
