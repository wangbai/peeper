package httpserv

import (
	"log"
	"net/http"
)

type Command interface {
	Handle(http.ResponseWriter, *http.Request)
}

var commands map[string]Command = make(map[string]Command)

func Register(name string, comm Command) {
	if comm == nil {
		log.Fatal("Command: ", name, " has nil handler")
	}

	if _, dup := commands[name]; dup {
		log.Fatal("Command: ", name, " has been registered")
	}

	commands[name] = comm
}

func matchCommand(name string) Command {
	comm, ok := commands[name]
	if ok {
		return comm
	}

	return defaultHandler{}
}
