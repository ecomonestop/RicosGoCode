package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//below enables parallelism as it says to use 2 threads.  now goroutine can run at the sametime, but is hard to account for all parallel issues when programming
	runtime.GOMAXPROCS(2)

	//using go keyword makes the func a go routine.  go routines run on top of threads communicating via channels.  go routines make it easier to do
	//concurrency.  go routines are good because they require less thread switching than normal concurrent programs.
	//concurrency is the composition of independently executing processes (processes saying lightly; independent more strictly speaking)
	//concurrency helps in DEALING with lots of things at once.  parallelism is about DOING lots of things at once
	var waitgroup sync.WaitGroup
	waitgroup.Add(2)

	go func() {
		defer waitgroup.Done()

		time.Sleep(5 * time.Second)
		fmt.Println("This is from first go routine")
	}()

	go func() {
		defer waitgroup.Done()
		fmt.Println("This is from second go routine")
	}()

	waitgroup.Wait()

}
