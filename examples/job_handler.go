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
		work := HeavyWork{Name: q, number: a}
		splanner.AddUnit(&work)
		_, _ = w.Write([]byte(fmt.Sprintf("job %s %d done", work.Name, a)))
	}
}

type HeavyWork struct {
	Name   string `json:"name"`
	number int
}

func (p *HeavyWork) Job() error {
	time.Sleep(500 * time.Millisecond)
	fmt.Println(fmt.Sprintf("heavy job is running %d", p.number))
	return nil
}
