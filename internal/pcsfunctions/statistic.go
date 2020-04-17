package pcsfunctions

import (
	"sync/atomic"
	"time"
	"xpan/baidupcs/expires"
)

// Statistic
type (
	Statistic struct {
		totalSize int64
		startTime time.Time
	}
)

// AddTotalSize
func (s *Statistic) AddTotalSize(size int64) int64 {
	return atomic.AddInt64(&s.totalSize, size)
}

// TotalSize
func (s *Statistic) TotalSize() int64 {
	return s.totalSize
}

// StartTimer
func (s *Statistic) StartTimer() {
	s.startTime = time.Now()
	expires.StripMono(&s.startTime)
}

// Elapsed
func (s *Statistic) Elapsed() time.Duration {
	return time.Now().Sub(s.startTime)
}
