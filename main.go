package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/tbarron-xyz/cider/structs"
	"net/http"
	// "strings"
	"encoding/json"
	"strconv"
)

var (
	STRINGS  = structs.STRINGS
	SETS     = structs.SETS
	LISTS    = structs.LISTS
	HASHES   = structs.HASHES
	COUNTERS = structs.COUNTERS
)

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

func single_message(cmd string) msi {
	// returns	{"status":"success", "response": something}
	// or 		{"status":"error", "response": err.Error()}
	var args []string
	var err error
	args, err = Parser(cmd)
	if err != nil {
		return errormsi(err)
	}
	if len(args) == 0 {
		return errormsi(fmt.Errorf("No arguments."))
	}
	handler, ok := arg0handlers[args[0]]
	if !ok {
		return errormsi(fmt.Errorf("Invalid command %s", args[0]))
	}
	var handled itf
	handled, err = handler.Handle(args[1:])
	if err != nil {
		return errormsi(err)
	} else {
		return successmsi(handled)
	}
}

func pipeline_message(cmd string) msi {
	var err error
	var cmds []string
	err = json.Unmarshal([]byte(cmd), &cmds)
	if err != nil {
		return errormsi(err)
	}
	single_responses := make([]msi, len(cmds))
	for i, e := range cmds {
		single_responses[i] = single_message(e)
	}
	return successmsi(single_responses)
}

func handle_message(cmd string) (tosend []byte) {
	verbose("<", cmd)
	var err error
	if len(cmd) == 0 {
		return
	}
	var handler func(string) msi
	if cmd[0] == '[' { // pipelining
		handler = pipeline_message
	} else { // single argument
		handler = single_message
	}
	tosend, err = json.MarshalIndent(handler(cmd), "", "    ") // json.Marshal(handler(cmd))
	if err != nil {
		panic(err.Error())
	}
	return
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}
		tosend := handle_message(string(p))
		err = conn.WriteMessage(messageType, tosend)
		if err != nil {
			return
		}
	}
}

var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/", wsHandler)
	fmt.Println("Listening on port", *_port)
	err := http.ListenAndServe(":"+strconv.Itoa(*_port), nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
