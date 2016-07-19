package config

import (
	"log"
)

type Buildup interface {
	ParseAndBuild(string)
}

var apps map[string]Buildup = make(map[string]Buildup)

func Register(name string, app Buildup) {
	if app == nil {
		log.Fatal("Application:%s has nil handler", name)
	}

	if _, dup := apps[name]; dup {
		log.Fatal("Application:%s has been registered", name)
	}

	apps[name] = app
}

func Build(dir string) {
	for _, a := range apps {
		a.ParseAndBuild(dir)
	}
}
