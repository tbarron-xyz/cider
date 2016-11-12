package structs

import (
	"sync"
	"fmt"
)

type StringsCollection struct {
	sync.RWMutex
	v msi 	// map[string]string
}

var STRINGS = &StringsCollection{v: msi{}}

func (this *StringsCollection) Set (key string, value itf) {
	// O(1)
	this.Lock()
		this.v[key] = value
	this.Unlock()
	return
}

func (this *StringsCollection) Get (key string) (response itf) {
	// O(1)
	this.RLock()
		response,_ = this.v[key]
	this.RUnlock()
	return
}

func (this *StringsCollection) Append (key, toappend string) (newval string, err error) {
	this.Lock()
		curval,_ := this.v[key]	// if doesn't exist, curval=""
		switch curval.(type) {
		case string:
			newval = curval.(string) + toappend
			this.v[key] = newval
		default:
			err = fmt.Errorf("%s is not a string", key)
		}
	this.Unlock()
	return
}

// func (this *StringsCollection) GetOrZero (key string) (response string) {
// 	this.RLock()
// 		value, ok := this.v[key]
// 		if !ok {
// 			this.RUnlock()
// 			this.Lock()
// 				this.v[key] = "0"
// 			this.Unlock()
// 			value = "0"
// 		}
// 		response = value
// 	this.RUnlock()
// 	return
// }