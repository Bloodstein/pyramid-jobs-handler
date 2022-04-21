package main

import (
	"log"
	"math/rand"
	"time"
)

type ProcessObject interface {
	Execute()
	Process()
}

type Process struct {
	Guid  string
	Abort chan bool
}

func (p Process) Process() {
	log.Printf("[ORDERING] [%s] I am working ...", p.Guid)
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
}

func (p Process) Execute() {
	for {
		select {
		case <-p.Abort:
			return
		default:
			p.Process()
		}
	}
}
