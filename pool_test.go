package hpool

import (
	"testing"
)

func TestPool(t *testing.T) {
	getCount := 10
	putCount := 20
	poolSize := 2

	p := NewPool(func() interface{} {
		return 123
	}, poolSize)
	p.SetDebug(true)

	for i := 0; i < getCount; i++ {
		p.Get()
	}
	for i := 0; i < putCount; i++ {
		p.Put(123)
	}
	pnewCount, pgetCount, pputCount, premoveCount := p.GetDebugInfo()
	if pnewCount != uint32(getCount) {
		t.Errorf("pool new count error.")
	}
	if pgetCount != uint32(getCount) {
		t.Errorf("pool get count error.")
	}
	if pputCount != uint32(putCount) {
		t.Errorf("pool put count error.")
	}
	if premoveCount != uint32(putCount-poolSize) {
		t.Errorf("pool remove count error.")
	}
}
