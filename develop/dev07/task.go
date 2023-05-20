package dev07

import (
	"fmt"
	"sync"
	"time"
)

func process(after time.Duration, data int) (out chan int) {
	ch := make(chan int)
	go func() {
		defer close(ch)
		time.Sleep(after)
		ch <- data
	}()
	return ch
}

func merge(channels ...<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		wg := sync.WaitGroup{}
		wg.Add(len(channels))
		for _, ch := range channels {
			go func(ch <-chan int) {
				for v := range ch {
					out <- v
				}
				wg.Done()
			}(ch)
		}
		wg.Wait()
		close(out)
	}()
	return out
}

func Main() {

	ch5 := process(time.Second*1, 1)
	ch4 := process(time.Second*2, 2)
	ch3 := process(time.Second*3, 3)
	ch2 := process(time.Second*4, 4)
	ch1 := process(time.Second*5, 5)

	res := merge(ch1, ch2, ch3, ch4, ch5)

	for v := range res {
		fmt.Println(v)
	}

}
