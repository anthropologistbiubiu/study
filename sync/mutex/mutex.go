package main

import (
	"runtime"
	"sync/atomic"
)

type Mutex struct {
	key  int32
	sema uint32
}

type Locker interface {
	Lock()
	Unlock()
}

func (m *Mutex) Lock() {
	if atomic.AddInt32(&m.key, 1) == 1 {
		return
	}
	runtime.Semacquire(&m.sema)
}

func (m *Mutex) Unlock() {
	switch v := atomic.AddInt32(&m.key, -1); {
	case v == 0:
		return
	case v == -1:
		panic("sync: unlock of unlocked mutex")
	}
	runtime.Semrelease(&m.sema)
}
