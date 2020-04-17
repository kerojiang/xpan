package cachepool

import (
	"runtime"
	"sync"
	"xpan/pcsutil/converter"
)

var (
	syncPoolSize     = int(64 * converter.KB)
	syncPoolFirstNew = false
	// SyncPool
	SyncPool = sync.Pool{
		New: func() interface{} {
			syncPoolFirstNew = true
			return RawMallocByteSlice(syncPoolSize)
		},
	}
)

// SetSyncPoolSize
func SetSyncPoolSize(size int) {
	if syncPoolFirstNew && size != syncPoolSize {
		runtime.GC()
	}
	syncPoolSize = size
}
