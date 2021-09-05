package abool

import "sync/atomic"

type AtomicBool uint32

// New creates AtomicBool with default set to false.
func New() *AtomicBool {
	return new(AtomicBool)
}

// Set changes the boolean to true
func (v *AtomicBool) Set() {
	atomic.StoreUint32((*uint32)(v), 1)
}

// Unset changes the boolean to false
func (v *AtomicBool) Unset() {
	atomic.StoreUint32((*uint32)(v), 0)
}

// Is returns whether the boolean is false or true
func (v *AtomicBool) Is() bool {
	return atomic.LoadUint32((*uint32)(v))&1 == 1
}
