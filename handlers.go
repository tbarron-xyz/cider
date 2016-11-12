package main

import (
	"fmt"
)

func nargchecker (cmd string, numargs, base, multiple int) (err error) {
	if numargs - base < 0 /*|| (numargs - base) % multiple != 0*/ {
		err = fmt.Errorf("Wrong number of arguments for %s: %d (expected: %d)", cmd, numargs, base)
		// err = fmt.Errorf("Wrong number of arguments for %s: %d (expected: %d+%dx)", cmd, numargs, base, multiple)
	}
	return
}

func argsplitter (args []string, multiple int) (splitargs [][]string) {
	if len(args) % multiple != 0 {
		return
	}
	for i,_ := range args {
		if i % multiple == 0 {
			splitargs = append(splitargs, args[i:i+multiple])
		}
	}
	return
}

func fakeItoa (s *bool, i int) {
	if i == 0 {
		*s = false
	} else if i == 1 {
		*s = true
	}
}





type cmdHandler struct {
	// expects (base + n*multiple) arguments, for some $n \in \NN_{\ge 0}$
	name string
	base int
	multiple int
	handle func ([]string) (itf, error)	// (string, error)
}

func (this *cmdHandler) Handle (args []string) (response itf, err error) {
	err = nargchecker(this.name, len(args), this.base, this.multiple)
	if err != nil {
		return
	}
	response, err = this.handle(args)
	return
}

func NewHandler (name string, base, multiple int, handle func ([]string) (itf, error)) (*cmdHandler) {
	handler := cmdHandler{name, base, multiple, handle}
	arg0handlers[name] = &handler
	return &handler
}

var arg0handlers = map[string]*cmdHandler{}

