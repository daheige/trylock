package trylock

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

const (
	LockedFlag   int32 = 1
	UnlockedFlag int32 = 0
)

type Mutex struct {
	in     sync.Mutex
	status *int32
}

func NewMutex() *Mutex {
	status := UnlockedFlag
	return &Mutex{
		status: &status,
	}
}

func (m *Mutex) Lock() {
	m.in.Lock()
}

func (m *Mutex) Unlock() {
	m.in.Unlock()
}

func (m *Mutex) TryLock() bool {
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.in)), UnlockedFlag, LockedFlag)
}

func (m *Mutex) IsLocked() bool {
	if atomic.LoadInt32(m.status) == LockedFlag {
		return true
	}
	return false
}

