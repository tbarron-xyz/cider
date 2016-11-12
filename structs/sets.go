package structs

import (
	"sync"
	// "fmt"
)

type set map[string]struct{}

type Set struct {
	sync.RWMutex
	s set	// underlying "s"et
}

type SetsCollection struct {
	sync.RWMutex
	v map[string]*Set
}

var SETS = &SetsCollection{v: map[string]*Set{}}

func NewEmptySet () (s *Set) {
	return &Set{s: set{}}
}

func (this *Set) Add (members ...string) (response int) {
	this.Lock()
		s := this.s
		for _,member := range members {
			_,ok := s[member]
			if !ok {
				s[member] = struct{}{}
				response += 1
			}
		}
	this.Unlock()
	return
}

func (this *Set) Remove (name string, members ...string) (response int) {
	this.Lock()
		for _,member := range members {
			_,ok := this.s[member]
			if ok {
				delete(this.s, member)
				response += 1
			}
		}
		if len(this.s) == 0 {
			defer SETS.Remove(name)
		}
	this.Unlock()
	return
}

func (this *Set) IsMember (member string) (response bool) {
	this.RLock()
		_,ok := this.s[member]
		if ok {
			response = true
		}
	this.RUnlock()
	return
}

func (this *Set) Members () (response []string) {
	this.RLock()
		s := this.s
		response = make([]string, len(s))
		i := 0
		for member,_ := range s {
			response[i] = member
		}
	this.RUnlock()
	return
}

func (this *Set) Intersect (others ...*Set) (response Set) {
	this.RLock()
		s := this.s
		if len(others) == 0 {
			return *this
		}
		response = Set{}
		for member,_ := range s {
			keep := true
			for _,keyset := range others {
				_,ok := keyset.s[member]
				if !ok {
					keep = false
					break
				}
			}
			if keep {
				response.s[member] = struct{}{}
			}
		}
	this.RUnlock()
	return
}

func (this *Set) Cardinality () (response int) {
	this.RLock()
		response = len(this.s)
	this.RUnlock()
	return response
}

func (this *Set) Diff (others ...*Set) (response *Set) {	// setminus
	this.RLock()
	for _,other := range others {
		other.RLock()
	}
		response = NewEmptySet()
		s := this.s
		// keysets := []Set{}
		// for _,key := range keys {
		// 	keyset, ok := SETS[key]
		// 	if ok {
		// 		keysets.append(keyset)
		// 	}
		// }
		for member,_ := range s {
			keep := true
			for _,other := range others {
				_,ok := other.s[member]
				if ok {
					keep = false
					break
				}
			}
			if keep {
				response.s[member] = struct{}{}
			}
		}
	this.RUnlock()
	for _,other := range others {
		other.RUnlock()
	}
	return
}

func (this *Set) Pop () (response string) {
	this.Lock()
		for member,_ := range this.s {	// random
			response = member
			delete(this.s, member)
			break
		}
	this.Unlock()
	return
}

func (this *Set) RandMember () (response string) {
	this.RLock()
		for member,_ := range this.s {
			response = member
			break
		}
	this.RUnlock()
	return
}

func (this *Set) Move (other *Set, member string) (moved bool) {
	// returns true if member in this and the move was performed, false if member not in this
	this.Lock()
		_, ok := this.s[member]
		if ok {
			other.Lock()
				other.s[member] = struct{}{}
			other.Unlock()
			delete(this.s, member)
			moved = true
		}
	this.Unlock()
	return
}







func (this *SetsCollection) Get (key string) (s *Set, ok bool) {
	this.RLock()
		s, ok = this.v[key]
	this.RUnlock()
	return
}

func (this *SetsCollection) GetOrCreateEmpty (key string) (s *Set) {
	var ok bool
	s, ok = this.Get(key)
	if !ok {
		this.Lock()
			s = NewEmptySet()	// pointer
			this.v[key] = s
		this.Unlock()
	}
	return
}

func (this *SetsCollection) Remove (key string) {
	this.Lock()
		delete(this.v, key)
	this.Unlock()
}