package viewmanager

import (
	"log"
)

type View interface {
	Update()
	View()
}

type ViewCreator func() View

var currentView View
var routes = make(map[string]ViewCreator)

func Open(route string) {
	create, ok := routes[route]
	if !ok {
		log.Fatal("call to nonexistent Route: '" + route + "'")
	}
	currentView = create()
}

func Get() View {
	return currentView
}

func Register(route string, fn ViewCreator) {
	routes[route] = fn
}
