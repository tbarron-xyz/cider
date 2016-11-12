package structs

import (
	"sync"
)

type IntsCollection struct {
	sync.RWMutex
	v map[string]int
}

var INTS = &IntsCollection{v: map[string]int{}}

func (this *IntsCollection) Set (key string, value int) (response string) {
	// O(1)
	this.Lock()
		this.v[key] = value
	this.Unlock()
	response = "OK"
	return
}

func (this *IntsCollection) Get (key string) (response int, ok bool) {
	// O(1)
	this.RLock()
		response, ok = this.v[key]
	this.RUnlock()
	return
}

func (this *IntsCollection) IncrBy (key string, incr int) (newval int) {
	// O(1)
	this.RLock()
		curval, ok := this.v[key]
	this.RUnlock()
	if !ok {
		curval = 0
	}
	newval = curval + incr
	this.Lock()
		this.v[key] = newval
	this.Unlock()
	return
}

func (this *IntsCollection) GetOrSetToZero (key string) (value int) {
	// O(1)
	var ok bool
	this.RLock()
		value, ok = this.v[key]
	this.RUnlock()
	if !ok {
		value = 0
		this.Lock()
			this.v[key] = 0
		this.Unlock()
	}
	return
}