package server

import (
	"log"
)

type Module interface {
	ParseAndLoad(dir string)
}

var modules map[string]Module = make(map[string]Module)

func RegisterModule(name string, m Module) {
	if m == nil {
		log.Fatal("Module:%s has nil handler", name)
	}

	if _, dup := modules[name]; dup {
		log.Fatal("Module:%s has been registered", name)
	}

	modules[name] = m
}

func LoadModule(dir string) {
	for _, m := range modules {
		m.ParseAndLoad(dir)
	}
}
