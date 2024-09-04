package topics

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AvoidingRace() {

	// Program to increment a variable by 1
	// var counter int
	// var wg sync.WaitGroup

	// for i := 0; i < 1000; i++ {
	// 	for i := 0; i < 10; i++ {
	// 		wg.Add(1)
	// 		go func(wg *sync.WaitGroup) {
	// 			defer wg.Done()
	// 			counter = counter + 1
	// 		}(&wg)
	// 	}
	// }

	// wg.Wait()
	// fmt.Println("Counter: ", counter)

	//Mutex
	solveByMutex()

	// atomicVariables
	// solveByAtomicVariables()

	// Channels
	solveByChannel()
}

func solveByMutex() {

	var counter int
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup, mutex *sync.Mutex) {
				defer wg.Done()
				mutex.Lock()
				counter = counter + 1
				mutex.Unlock()
			}(&wg, &mutex)
		}
	}
	wg.Wait()
	fmt.Println("Counter using mutex: ", counter)

}

func solveByAtomicVariables() {

	var counter int64
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				// counter = counter + 1
				atomic.AddInt64(&counter, 1)
			}(&wg)
		}
	}
	wg.Wait()
	fmt.Println("Counter using atomic variables: ", counter)

}

func solveByChannel() {

	var counter int
	ch := make(chan int)
	var wg sync.WaitGroup

	go func() {
		for data := range ch {
			counter = counter + data
		}
	}()

	for i := 0; i < 1000; i++ {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				ch <- 1
			}(&wg)
		}
	}

	wg.Wait()
	close(ch)
	fmt.Println("Counter using channel: ", counter)

}
