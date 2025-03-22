// problem in go routine , we have to write sleep in order to wait
// for execution

package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	// defer runs at the end same as cleanup func
	defer w.Done()
	fmt.Println("doing task", id)

}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()
}
