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

	// init the dispatcher & keep it listening.
	splanner.NewDispatcher(15).Run(true)

	s := http.Server{}
	s.Addr = ":8080"
	addJobRoute()
	log.Fatal(s.ListenAndServe())
}
