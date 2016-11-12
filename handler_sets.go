package main

import (
	// "fmt"
)

func init () {
	NewHandler(	// SADD key member [member ...]
		"SADD", 2, 1, func (args []string) (response itf, err error) {
			key, values := args[0], args[1:]
			s := SETS.GetOrCreateEmpty(key)
			response = s.Add(values...)
			return
		})

	NewHandler(	// SREM key member [member ...]
		"SREM", 2, 1, func (args []string) (response itf, err error) {
			key, values := args[0], args[1:]
			s, ok := SETS.Get(key)
			if ok {
				response = s.Remove(key, values...)
			} else {
				response = 0
			}
			return
		})

	NewHandler(
		"SISMEMBER", 2, 0, func (args []string) (response itf, err error) {
			key, value := args[0], args[1]
			s, ok := SETS.Get(key)
			if ok {
				response = s.IsMember(value)
			}
			return
		})

	NewHandler(
		"SMEMBERS", 1, 0, func (args []string) (response itf, err error) {
			key := args[0]
			s, ok := SETS.Get(key)
			if ok {
				_ = s
				response = ""	// CODE THIS
			}
			//  else {
			// 	err = fmt.Errorf("Set does not exist.")
			// }
			return
		})

	NewHandler(
		"SCARD", 1, 1, func (args []string) (response itf, err error) {
			key := args[0]
			s, ok := SETS.Get(key)
			if ok {
				response = s.Cardinality()
			}
			//  else {
			// 	err = fmt.Errorf("Set does not exist.")
			// }
			return
		})

	NewHandler(
		"SPOP", 1, 0, func (args []string) (response itf, err error) {
			key := args[0]
			s, ok := SETS.Get(key)
			if ok {
				response = s.Pop()
			}
			//  else {
			// 	err = fmt.Errorf("Set does not exist.")
			// }
			return
		})

	NewHandler(
		"SRANDMEMBER", 1, 0, func (args []string) (response itf, err error) {
			key := args[0]
			s, ok := SETS.Get(key)
			if ok {
				response = s.RandMember()
			}
			// else {
			// 	err = fmt.Errorf("Set does not exist.")
			// }
			return
		})
}