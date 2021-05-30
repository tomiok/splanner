package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

type jobHandler struct {
}

func (j *jobHandler) JobHandler(ctx *fiber.Ctx) error {
	payload := Payload{Name: ctx.Get("q", "default")}
	job := Job{Payload: payload}
	sp := Unit{job: job.Payload.Run}
	jobQueue <- sp
	_, err := ctx.Write([]byte("job " + job.Payload.Name + " done"))
	return err
}

type Job struct {
	Payload Payload
}

type Payload struct {
	Name string `json:"name"`
}

func (p *Payload) Run() {
	time.Sleep(2 * time.Second)
	fmt.Println("heavy jobHandler is running")
}
