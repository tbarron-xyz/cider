package structs

import (
	"sync"
	"time"
)

type Counter struct {
	sync.RWMutex
	value    int
	duration int // milliseconds
	timerId int
}

func NewCounter() (c *Counter) {
	return &Counter{duration: 5000}
}

type CountersCollection struct {
	sync.RWMutex
	v map[string]*Counter
}

func (this *CountersCollection) Get(key string) (c *Counter, ok bool) {
	this.Lock()
	c, ok = this.v[key]
	this.Unlock()
	return
}

func (this *CountersCollection) GetOrInit(key string) (c *Counter) {
	var ok bool
	c, ok = this.Get(key)
	if !ok {
		this.Lock()
		c = NewCounter()
		this.v[key] = c
		this.Unlock()
	}
	return
}

var COUNTERS = &CountersCollection{v: map[string]*Counter{}}

func (this *Counter) Get() (value int) {
	this.RLock()
	value = this.value
	this.RUnlock()
	return
}


func (this *Counter) incrClosure(incr, timerId int) (closure func()) {
	closure = func() {
		this.Lock()
		if timerId == this.timerId {
			this.value += incr
		}
		this.Unlock()
	}
	return
}

func (this *Counter) SetDuration(duration int) {
	this.Lock()
	this.duration = duration
	this.Unlock()
	return
}

func (this *Counter) IncrBy(incr int) (newval int) {
	this.Lock()
	this.value += incr
	newval = this.value
	duration := this.duration
	timerId := this.timerId
	this.Unlock()
	time.AfterFunc(time.Duration(duration)*time.Millisecond, this.incrClosure(-incr,timerId))
	return
}

func (this *Counter) Reset() {
	this.Lock()
	this.value = 0
	this.timerId += 1
	this.Unlock()
}
