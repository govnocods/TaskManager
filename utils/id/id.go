package id

import "sync/atomic"

var counter int64

func NextID() int64 {
	return atomic.AddInt64(&counter, 1)
}
