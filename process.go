package gotp

import "log"

var logger *log.Logger

type Message[S any] interface {
	Apply(S) S
}

type Process[S any] struct {
	PID   int
	State S
	In    chan Message[S]
}

func NewProcess[S any](state S) *Process[S] {
	PID := 1
	p := Process[S]{
		PID:   PID,
		State: state,
		In:    make(chan Message[S]),
	}
	go p.Run()
	return &p
}

func (p *Process[S]) Run() {
	logger.Println("Running Process")
	for msg := range p.In {
		logger.Println("Received message")
		newState := msg.Apply(p.State)
		logger.Printf("NewState => %+v \n", newState)
		p.State = newState
	}
}
