package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

type jobHandler struct {
}

func (j *jobHandler) JobHandler(ctx *fiber.Ctx) error {
	payload := Payload{S: ctx.Get("q", "default")}
	job := Job{P: payload}
	jobQueue <- job
	return nil
}

type Job struct {
	P Payload
}

type Payload struct {
	S string `json:"s"`
}

func (p *Payload) Run() {
	time.Sleep(2 * time.Second)
	fmt.Println("heavy jobHandler is running")
}
