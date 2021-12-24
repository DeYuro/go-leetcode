package main

import (
	"fmt"
	"runtime"
)
// infinity loop
func main() {
	runtime.GOMAXPROCS(1) // > 1 for finished

	done := false

	go func() {
		done = true
	}()

	for !done {
	}
	fmt.Println("finished")
}
