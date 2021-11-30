package main

import (
	"fmt"
	"github.com/tomiok/splanner"
	"log"
	"net/http"
	"time"
)

func JobHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("getting job...")
	q := r.URL.Query().Get("q")

	if q == "" {
		q = "default"
	}

	for a := 0; a < 100; a++ {
		payload := Payload{Name: q, number: a}
		job := Job{Payload: payload}
		workUnit := splanner.Unit{Job: job.Payload.Run}
		splanner.JobQueue <- workUnit
		_, _ = w.Write([]byte(fmt.Sprintf("job %s %d done", payload.Name, a)))
	}
}

type Job struct {
	Payload Payload
}

type Payload struct {
	Name   string `json:"name"`
	number int
}

func (p *Payload) Run() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println(fmt.Sprintf("heavy job is running %d", p.number))
}
