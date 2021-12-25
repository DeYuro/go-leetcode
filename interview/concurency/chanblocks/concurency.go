package main

import "time"

func main() {
	timeStart := time.Now()

	// for 3 sec instead of 	_, _ = <-worker(), <-worker()
	//ch1, ch2 := worker(), worker()
	//_, _ = <-ch1, <-ch2
	_, _ = <-worker(), <-worker()

	println(int(time.Since(timeStart).Seconds())) // 3 or 6
}
// 6 sec coz worker block main

func worker() chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 42
	}()

	return ch
}