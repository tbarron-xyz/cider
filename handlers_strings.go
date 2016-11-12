package main

import (

)

func init () {
	NewHandler(	// todo: handle e.g. `SET field1 value1 field2 value2`
		"SET", 2, 2, func (args []string) (response itf, err error) {	// cannot fail; returns an empty string
			key, value := args[0], args[1]
			STRINGS.Set(key, value)
			return
		})

	NewHandler(
		"GET", 1, 1, func (args []string) (response itf, err error) {	// returns empty string if value does not exist
			key := args[0]
			response = STRINGS.Get(key)
			return
		})

	NewHandler(
		"APPEND", 2, 2, func (args []string) (response itf, err error) {	// returns the new value of the string
			var keyvalues [][2]string
			for i,_ := range args {
				if i % 2 == 0 {
					keyvalues = append(keyvalues, [2]string{args[i], args[i+1]})
				}
			}
			for _,e := range keyvalues {
				key, value := e[0], e[1]
				STRINGS.Append(key, value)
			}
			return
		})
}