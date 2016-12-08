package main

import (
	"strconv"
	"unicode"
)

func init() {
	NewHandler(
		"HSET", 3, func(args []string) (response itf, err error) {
			var value itf
			var key, field string
			key, field, value = args[0], args[1], args[2]
			h := HASHES.GetOrInit(key)
			var isint = true
			for _, e := range args[2] {
				if !unicode.IsNumber(e) {
					isint = false
				}
			}
			if isint {
				value, err = strconv.Atoi(args[4])
				if err != nil {
					return
				}
			}
			response = h.Set(field, value)
			return
		})

	NewHandler(
		"HGET", 2, func(args []string) (response itf, err error) {
			key, field := args[0], args[1]
			h, ok := HASHES.Get(key)
			if ok {
				response, ok = h.Get(field)
			}
			return
		})

	NewHandler(
		"HLEN", 1, func(args []string) (response itf, err error) {
			key := args[0]
			h, ok := HASHES.Get(key)
			if ok {
				response = h.Len()
			}
			return
		})

	NewHandler(
		"HKEYS", 1, func(args []string) (response itf, err error) {
			key := args[0]
			h, ok := HASHES.Get(key)
			if ok {
				response = h.Keys()
			}
			return
		})

	// NewHandler(
	// 	"HINCRBY", 3, func(args []string) (response itf, err error) {
	// 		key, field, incr := args[0], args[1], args[2]
	// 		h := HASHES.GetOrInit(key)
	// 		var incrint int
	// 		incrint, err = strconv.Atoi(incr)
	// 		if err != nil {
	// 			return
	// 		}
	// 		response, err = h.IncrBy(field, incrint)
	// 		return
	// 	})
}
