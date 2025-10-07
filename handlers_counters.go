package main

import (
	"strconv"
)

func init() {
	NewHandler(
		"INCRBY", 2, func(args []string) (response itf, err error) {
			key, value := args[0], args[1]
			c := server.Counters.GetOrInit(key)
			var i int
			i, err = strconv.Atoi(value)
			if err != nil {
				return
			}
			response = c.IncrBy(i)
			return
		})

	NewHandler(
		"CGET", 1, func(args []string) (response itf, err error) {
			key := args[0]
			c, ok := server.Counters.Get(key)
			if ok {
				response = c.Get()
			} else {
				response = 0
			}
			return
		})

	NewHandler(
		"DURATION", 2, func(args []string) (response itf, err error) {
			key, value := args[0], args[1]
			c, ok := server.Counters.Get(key)
			if ok {
				var i int
				i, err = strconv.Atoi(value)
				if err != nil {
					return
				}
				c.SetDuration(i)
			}
			return
		})

	NewHandler(
		"CRESET", 1, func(args []string) (response itf, err error) {
			key := args[0]
			c, ok := server.Counters.Get(key)
			if ok {
				c.Reset()
			}
			return
		})
}
