package gotp

import (
	"log"
	"testing"
	"time"
)

type State struct {
	counter int
}

type SayHello struct {
	Name string
}

func (sh SayHello) Apply(st State) State {
	logger.Printf("Updating state to %v \n", st.counter+1)
	return State{counter: st.counter + 1}
}

func Test_NewProcess(t *testing.T) {
	logger = log.Default()

	actor := NewProcess(State{counter: 0})
	actor.In <- SayHello{}
	time.Sleep(5 * time.Second)

	logger.Printf("actor.State.counter => %d \n", actor.State.counter)

	if actor.State.counter != 1 {
		t.Errorf("State not updated")
	}
}
