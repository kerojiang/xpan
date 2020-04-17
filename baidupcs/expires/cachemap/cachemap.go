package cachemap

import (
	"sync"
	"xpan/baidupcs/expires"
)

var (
	// GlobalCacheOpMap
	GlobalCacheOpMap = CacheOpMap{}
)

// CacheOpMap
type (
	CacheOpMap struct {
		cachePool sync.Map
	}
)

// LazyInitCachePoolOp
func (cm *CacheOpMap) LazyInitCachePoolOp(op string) CacheUnit {
	cacheItf, _ := cm.cachePool.LoadOrStore(op, &cacheUnit{})
	return cacheItf.(CacheUnit)
}

// RemoveCachePoolOp
func (cm *CacheOpMap) RemoveCachePoolOp(op string) {
	cm.cachePool.Delete(op)
}

// ClearInvalidate 清除已过期的数据(一般用不到)
func (cm *CacheOpMap) ClearInvalidate() {
	cm.cachePool.Range(func(_, cacheItf interface{}) bool {
		cache := cacheItf.(CacheUnit)
		cache.Range(func(key interface{}, exp expires.DataExpires) bool {
			if exp.IsExpires() {
				cache.Delete(key)
			}
			return true
		})
		return true
	})
}

// PrintAll 输出所有缓冲项目
func (cm *CacheOpMap) PrintAll() {
}
