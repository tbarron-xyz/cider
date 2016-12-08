package main

import "fmt"

func nargchecker(cmd string, numargs, base, multiple int) (err error) {
	if numargs-base < 0 /*|| (numargs - base) % multiple != 0*/ {
		err = fmt.Errorf("Wrong number of arguments for %s: %d (expected: %d)", cmd, numargs, base)
		// err = fmt.Errorf("Wrong number of arguments for %s: %d (expected: %d+%dx)", cmd, numargs, base, multiple)
	}
	return
}

func split_args(args []string, multiple int) (splitargs [][]string) {
	if len(args)%multiple != 0 {
		return
	}
	for i, _ := range args {
		if i%multiple == 0 {
			splitargs = append(splitargs, args[i:i+multiple])
		}
	}
	return
}

type itf interface{}
type msi map[string]itf

func errormsi(err error) msi {
	return msi{"status": "error", "error": err.Error()}
}

func successmsi(response interface{}) msi {
	return msi{"status": "success", "response": response}
}

func verbose(args ...interface{}) {
	if *_verbose {
		fmt.Println(args...)
	}
}
