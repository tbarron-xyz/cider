package main

func init() {
	NewHandler( // SADD key member [member ...]
		"SADD", 2, func(args []string) (response itf, err error) {
			key, values := args[0], args[1:]
			s := SETS.GetOrCreateEmpty(key)
			response = s.Add(values...)
			return
		})

	NewHandler( // SREM key member [member ...]
		"SREM", 2, func(args []string) (response itf, err error) {
			key, values := args[0], args[1:]
			s, ok := SETS.Get(key)
			if ok {
				response = s.Remove(values...)
			} else {
				response = 0
			}
			return
		})

	NewHandler(
		"SISMEMBER", 2, func(args []string) (response itf, err error) {
			key, value := args[0], args[1]
			s, ok := SETS.Get(key)
			if ok {
				response = s.IsMember(value)
			}
			return
		})

	NewHandler(
		"SMEMBERS", 1, func(args []string) (response itf, err error) {
			key := args[0]
			s, ok := SETS.Get(key)
			if ok {
				_ = s
				response = "" // CODE THIS
			}
			return
		})

	NewHandler(
		"SCARD", 1, func(args []string) (response itf, err error) {
			key := args[0]
			s, ok := SETS.Get(key)
			if ok {
				response = s.Cardinality()
			}
			return
		})

	NewHandler(
		"SPOP", 1, func(args []string) (response itf, err error) {
			key := args[0]
			s, ok := SETS.Get(key)
			if ok {
				response = s.Pop()
			}
			return
		})

	NewHandler(
		"SRANDMEMBER", 1, func(args []string) (response itf, err error) {
			key := args[0]
			s, ok := SETS.Get(key)
			if ok {
				response = s.RandMember()
			}
			return
		})
}
