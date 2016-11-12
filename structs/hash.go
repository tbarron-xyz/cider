package structs

import (
	// "strconv"
	"fmt"
	"sync"
)

func fff() {
	fmt.Println("fff")
}

type itf interface{}
type msi map[string]itf

type hash msi // map[string]string

type Hash struct {
	sync.RWMutex
	h hash
}

type HashesCollection struct {
	sync.RWMutex
	v map[string]*Hash
}

func NewEmptyHash() *Hash {
	return &Hash{h: hash{}}
}

var HASHES = &HashesCollection{v: map[string]*Hash{}}

func (this *Hash) Exists(field string) (response int) {
	// O(1)
	this.RLock()
	_, ok := this.h[field]
	if ok {
		response = 1
	}
	this.RUnlock()
	return
}

func (this *Hash) Keys() (keys []string) {
	// O(len(this))
	this.RLock()
	h := this.h
	keys = make([]string, len(h))
	i := 0
	for key := range h {
		keys[i] = key
		i++
	}
	this.RUnlock()
	return
}

func (this *Hash) Vals() (values []itf) {
	// O(len(this))
	this.RLock()
	h := this.h
	values = make([]itf, len(h))
	i := 0
	for _, value := range h {
		values[i] = value
		i++
	}
	this.RUnlock()
	return
}

func (this *Hash) Get(field string) (value itf, ok bool) {
	// O(1)
	this.RLock()
	value, ok = this.h[field]
	this.RUnlock()
	return
}

func (this *Hash) HGETALL() (eh Hash) {
	// O(len(this))
	return *this
}

func (this *Hash) Set(field string, value itf) (response string) {
	// returns 1 if new field, 0 if updated existing field
	// O(1)
	this.Lock()
	_, ok := this.h[field]
	if !ok {
		response = "1"
	} else {
		response = "0"
	}
	this.h[field] = value
	this.Unlock()
	return
}

func (this *Hash) Del(fields []string) (response int) {
	// O(1)
	this.Lock()
	h := this.h
	for _, field := range fields {
		_, ok := h[field]
		if ok {
			delete(h, field)
			response += 1
		}
	}
	this.Unlock()
	return
}

func (this *Hash) Len() (length int) {
	// O(1)
	this.RLock()
	length = len(this.h)
	this.RUnlock()
	return
}

func (this *Hash) IncrBy(field string, incr int) (value int, err error) {
	// O(1)
	this.Lock()
	h := this.h
	curval, ok := h[field]
	if !ok {
		value = 1
		h[field] = "1"
		this.Unlock()
		return
	}
	switch curval.(type) {
	case int:
		value = curval.(int) + incr
		h[field] = value
	default:
		err = fmt.Errorf("Field value is not an int.")
	}
	// var intval int
	// intval, err = strconv.Atoi(curval)
	// if err != nil { return }
	// value = intval + incr
	// h[field] = strconv.Itoa(value)
	this.Unlock()
	return
}

func (this *HashesCollection) GetOrInit(key string) (response *Hash) {
	var ok bool
	response, ok = this.Get(key)
	if !ok {
		response = NewEmptyHash()
		this.Lock()
		this.v[key] = response
		this.Unlock()
	}
	return
}

func (this *HashesCollection) Get(key string) (response *Hash, ok bool) {
	// O(1)
	this.RLock()
	response, ok = this.v[key]
	this.RUnlock()
	return
}
