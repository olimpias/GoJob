package GoJob

import "sync"

type AtomicBool struct {
	boolValue bool;
	mutex sync.RWMutex;
}

func NewAtomicBool() AtomicBool  {
	return AtomicBool{boolValue:false};
}

func NewAtomicBoolWithPtr()  *AtomicBool{
	return &AtomicBool{boolValue:false};
}

func (value * AtomicBool) getBoolValue() bool  {
	value.mutex.RLock();
	defer value.mutex.RUnlock();
	return value.boolValue;
}

func (value * AtomicBool) setBoolValue(newValue bool)  {
	value.mutex.Lock();
	defer value.mutex.Unlock();
	value.boolValue = newValue;
}
