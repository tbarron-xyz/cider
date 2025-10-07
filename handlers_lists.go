package main

func init() {
	NewHandler( // todo: handle e.g. `SET field1 value1 field2 value2`
		"LPUSH", 2, func(args []string) (response itf, err error) { // returns the new length of the list
			key, value := args[0], args[1]
			l := server.Lists.GetOrInit(key)
			response = l.LPush(value)
			return
		})

	NewHandler( // todo: handle e.g. `SET field1 value1 field2 value2`
		"RPUSH", 2, func(args []string) (response itf, err error) { // returns the new length of the list
			key, value := args[0], args[1]
			l := server.Lists.GetOrInit(key)
			response = l.RPush(value)
			return
		})

	NewHandler( // todo: handle e.g. `SET field1 value1 field2 value2`
		"LPOP", 1, func(args []string) (response itf, err error) { // returns the new length of the list
			key := args[0]
			l, ok := server.Lists.Get(key)
			if ok {
				response, _ = l.LPop()
			}
			return
		})

	NewHandler( // todo: handle e.g. `SET field1 value1 field2 value2`
		"RPOP", 1, func(args []string) (response itf, err error) { // returns the new length of the list
			key := args[0]
			l, ok := server.Lists.Get(key)
			if ok {
				response, _ = l.RPop()
			}
			return
		})
	NewHandler(
		"LLEN", 1, func(args []string) (response itf, err error) {
			key := args[0]
			l, ok := server.Lists.Get(key)
			if ok {
				response = l.Len()
			}
			return
		})

}
