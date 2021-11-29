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

	payload := Payload{Name: q}
	job := Job{Payload: payload}
	sp := splanner.Unit{Job: job.Payload.Run}
	splanner.JobQueue <- sp
	_, _ = w.Write([]byte("job " + job.Payload.Name + " done"))
}

type Job struct {
	Payload Payload
}

type Payload struct {
	Name string `json:"name"`
}

func (p *Payload) Run() {
	time.Sleep(2 * time.Second)
	fmt.Println("heavy job is running...")
}
