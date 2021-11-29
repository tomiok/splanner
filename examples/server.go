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
	splanner.InitQueue()

	dispatcher := splanner.NewDispatcher()
	dispatcher.Run()

	s := http.Server{}
	s.Addr = ":8080"
	addJobRoute()
	log.Fatal(s.ListenAndServe())
}
