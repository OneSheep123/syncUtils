package utils

import (
	"github.com/petermattis/goid"
	"sync"
	"sync/atomic"
)

// ReentrantLock 可重入锁
type ReentrantLock struct {
	sync.Mutex
	owner     int64
	recursion int32
}

func (m *ReentrantLock) Lock() {
	gid := goid.Get()
	if atomic.LoadInt64(&m.owner) == gid {
		m.recursion++
		return
	}
	m.Mutex.Lock()
	atomic.StoreInt64(&m.owner, gid)
	m.recursion = 1
}

func (m *ReentrantLock) UnLock() {
	gid := goid.Get()
	if atomic.LoadInt64(&m.owner) == gid {
		panic("当前锁不允许被解锁")
	}
	m.recursion--
	if m.recursion != 0 {
		return
	}
	atomic.StoreInt64(&m.owner, -1)
	m.Unlock()
}
