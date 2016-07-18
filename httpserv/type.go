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
		log.Fatal("Command:%s has nil handler", name)
	}

	if _, dup := commands[name]; dup {
		log.Fatal("Command:%s has been registered", name)
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
