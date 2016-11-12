package structs

import (
	"sync"
	"time"
)

type Counter struct {
	sync.RWMutex
	value    int
	duration int // milliseconds
}

type CountersCollection struct {
	sync.RWMutex
	v map[string]*Counter
}

var COUNTERS = &CountersCollection{v: map[string]*Counter{}}

func (this *Counter) Get() (value int) {
	this.RLock()
	value = this.value
	this.RUnlock()
	return
}

func (this *Counter) incrBy(incr int) (newval int) {

	return
}

func (this *Counter) decrClosure(incr int) (closure func()) {
	closure = func() {
		this.Lock()
		this.value += incr
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
	this.Unlock()
	time.AfterFunc(time.Duration(duration)*time.Millisecond, this.decrClosure(incr))
	return
}
