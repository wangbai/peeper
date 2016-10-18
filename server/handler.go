package server

import (
	"log"
	"net/http"
)

var handlers map[string]http.Handler = make(map[string]http.Handler)

func RegisterHandler(name string, h http.Handler) {
	if h == nil {
		log.Fatal("Service:", name, " has nil handler")
	}

	if _, dup := handlers[name]; dup {
		log.Fatal("Service:", name, " has been registered")
	}

	handlers[name] = h
}

func matchHandler(name string) http.Handler {
	h, ok := handlers[name]
	if ok {
		return h
	}

	return defaultHandler{}
}
