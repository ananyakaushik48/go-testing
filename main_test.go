package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestState(t *testing.T) {
	state := State{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// doing work and keeping track of the task in waitgrp
		go func(i int) {
			state.count = i + 1
			wg.Done() // done working updating state on waitgrp
		}(i)
	}
	wg.Wait()
	fmt.Printf("%+v\n", state)
}