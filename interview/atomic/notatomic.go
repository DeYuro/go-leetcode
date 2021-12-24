package main

// Problem not atomic counter, no wg
func main() {
	var counter int
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
			println(counter)
		}()
	}
}

// Solve
//func main() {
//	var counter int32
//	var wg sync.WaitGroup
//	for i := 0; i < 1000; i++ {
//		wg.Add(1)
//		go func() {
//			atomic.AddInt32(&counter, 1)
//			wg.Done()
//		}()
//	}
//
//	wg.Wait()
//}