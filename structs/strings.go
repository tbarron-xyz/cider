package structs

import (
	"sync"
)

type StringsCollection struct {
	sync.RWMutex
	v map[string]string // map[string]string
}

var STRINGS = &StringsCollection{v: map[string]string{}}

func (this *StringsCollection) Set(key string, value string) {
	// O(1)
	this.Lock()
	this.v[key] = value
	this.Unlock()
	return
}

func (this *StringsCollection) Get(key string) (response itf) {
	// O(1)
	this.RLock()
	response, _ = this.v[key]
	this.RUnlock()
	return
}

func (this *StringsCollection) Append(key, toappend string) (newval string) {
	this.Lock()
	curval, _ := this.v[key] // if doesn't exist, curval=""
	newval = curval + toappend
	this.v[key] = newval
	this.Unlock()
	return
}
