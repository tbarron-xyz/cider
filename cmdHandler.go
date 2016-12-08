package main

// Do not construct this directly. Use NewHandler().
type cmdHandler struct {
	// expects (base + n*multiple) arguments, for some $n \in \NN_{\ge 0}$
	name   string
	base   int
	handle func([]string) (itf, error) // (string, error)
}

func (this *cmdHandler) Handle(args []string) (response itf, err error) {
	err = nargchecker(this.name, len(args), this.base, 0)
	if err != nil {
		return
	}
	response, err = this.handle(args)
	return
}

func NewHandler(name string, base int, handle func([]string) (itf, error)) *cmdHandler {
	handler := cmdHandler{name, base, handle}
	arg0handlers[name] = &handler
	return &handler
}

// GET, SET, INCRBY, etc.
var arg0handlers = map[string]*cmdHandler{}
