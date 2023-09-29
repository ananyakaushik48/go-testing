package main

import (
	"fmt"
	"sync"
)

// state manipulation example
type State struct {
	/*
		    Mutex -> Mutually Exclusive:
			A mutex is like a Pass at an event.
			At the event there are 100 speakers, but all the speakers speaking at once will fuck the scene up.
			So the Event moderator gives a flag to the speaker who is ready first. All the speakers have to hold the flag to speak.
			The speakers without the flag can only inform the Event Moderator about their readyness and accept the flag when the moderator hands it to them.

			In this case the speaker is the thread, the speaking is the resource and the Flag is the mutex
	*/
	sync.Mutex
	count int
}

func (s *State) setState(i int) {
	s.count = i
}

func main() {
}
