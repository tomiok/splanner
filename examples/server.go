package main

import (
	"github.com/tomiok/splanner"
	"log"
	"net/http"
)

func addJobRoute() {
	http.HandleFunc("/jobs", JobHandler)
}

func main() {
	// init the shared data structure.
	splanner.InitQueue(20)

	// init the dispatcher
	dispatcher := splanner.NewDispatcher(15)

	// the dispatcher is listening
	dispatcher.Run()

	s := http.Server{}
	s.Addr = ":8080"
	addJobRoute()
	log.Fatal(s.ListenAndServe())
}
