package structs

import (
	"fmt"
	"sync"
)

type listElement struct {
	prev *listElement
	value string
	next *listElement
}

type List struct {
	sync.RWMutex
	c chan struct{}
	head *listElement
	length int
	tail *listElement
}

type ListsContainer struct {
	sync.RWMutex
	v map[string]*List
}

func NewEmptyList () (l *List) {
	return &List{c: make(chan struct{})}
}

var LISTS = &ListsContainer{v: map[string]*List{}}


func (this *List) lpop () (value string) {
	value = this.head.value
	this.head = this.head.next
	if this.head != nil {
		this.head.prev = nil
	}
	this.length -= 1
	return
}

func (this *List) LPop () (value string, ok bool) {
	// O(1)
	this.RLock()
		length := this.length
	this.RUnlock()
	if length == 0 {
		return
	}
	this.Lock()
		value = this.lpop()
	this.Unlock()
	ok = true
	return
}

func (this *List) BLpop (timeout int) (value string, ok bool) {
	// O(1)
	this.RLock()
	if len(this.c) == 0 {
		this.RUnlock()
		<- this.c
	} else {
		this.RUnlock()
	}
	this.Lock()
	return
}

func (this *List) rpop () (value string) {
	value = this.tail.value
	this.tail = this.tail.prev
	if this.tail != nil {
		this.tail.next = nil
	}
	this.length -= 1
	return
}

func (this *List) RPop () (value string, ok bool) {
	// O(1)
	this.RLock()
		length := this.length
	this.RUnlock()
	if length == 0 {
		return
	}
	this.Lock()
		if this != nil {
		}
		value = this.rpop()
	this.Unlock()
	ok = true
	return
}

func (this *List) LPush (arg string) (length int) {	// returns the new length of the list
	// O(1)
	this.Lock()
		this.head = &listElement{prev: nil, value: arg, next: this.head}
		if this.head.next != nil {
			this.head.next.prev = this.head
		} else {
			this.tail = this.head
		}
		this.length += 1
		length = this.length
	this.Unlock()
	// this.c <- struct{}{}
	return
}

func (this *List) RPush (arg string) (length int) {
	// O(1)
	this.Lock()
		this.tail = &listElement{prev: this.tail, value: arg, next: nil}
		if this.tail.prev != nil {
			this.tail.prev.next = this.tail
		} else {	// list was empty
			this.head = this.tail
		}
		this.length += 1
		length = this.length
	this.Unlock()
	// this.c <- struct{}{}
	return
}

func (this *List) Set (index int, value string) (err error) {
	// O(1)
	cur := this.head
	for i:=0; i<index; i++ {
		cur = cur.next
		if cur == nil {
			return fmt.Errorf("Index out of range.")
		}
	}
	cur.value = value
	return
}

func (this *List) Len () int {
	// O(1)
	return this.length
}






func (this *ListsContainer) Get (key string) (l *List, ok bool) {
	this.RLock()
		l, ok = this.v[key]
	this.RUnlock()
	return
}

func (this *ListsContainer) GetOrInit (key string) (l *List) {
	l, ok := this.Get(key)
	if !ok {
		this.Lock()
			l = NewEmptyList()
			this.v[key] = l
		this.Unlock()
	}
	return
}