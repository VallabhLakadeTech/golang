package topics

import (
	"fmt"
	"sync"
)

func Workerpool() {

	workers := 3

	urlList := []int{1, 2, 3, 4, 5}

	inputCh := make(chan int, 2)
	outputCh := make(chan int, 2)

	var wg sync.WaitGroup
	wg.Add(len(urlList))

	for i := 0; i < workers; i++ {
		go worker(inputCh, outputCh, i, &wg)
	}

	for i := 0; i < len(urlList); i++ {
		inputCh <- urlList[i]
	}
	close(inputCh)

	go func() {
		wg.Wait()
		close(outputCh)
	}()

	for outputAggregator := range outputCh {
		fmt.Println("Response received from URL", outputAggregator)
	}

}

// inputCh is the read channel and outputCh is the write channel
func worker(inputCh <-chan int, outputCh chan<- int, workerId int, wg *sync.WaitGroup) {

	for data := range inputCh {
		// fmt.Println("Data:", data)
		outputCh <- callExternalAPI(data, workerId)
		wg.Done()
	}

}

func callExternalAPI(data, workerID int) int {
	fmt.Printf("URL %v is handled by worker %v\n", data, workerID)
	return data
}
