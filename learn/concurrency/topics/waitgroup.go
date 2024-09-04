package topics

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroup() {
	var wg sync.WaitGroup
	wg.Add(2)
	go count("Sheep", &wg)
	go count("Fox", &wg)
	wg.Wait()
}
func count(msg string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 200)
	}
	wg.Done()
}
