## gowatch

Simple watcher for TCP or UDP network service.

## Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/terryh/gowatch"
)

func main() {
	watchDog1, err := gowatch.NewWatcher("tcp", "www.google.com.tw:80", time.Second*5)
	watchDog2, err := gowatch.NewWatcher("tcp", "localhost:3000", time.Second*1)

	fmt.Println(watchDog1, err)
	fmt.Println(watchDog2, err)

	//select {}
	for {

		select {

		case msg := <-watchDog1.Status:
			fmt.Println("watchDog1", msg)
		case msg := <-watchDog2.Status:
			fmt.Println("watchDog2", msg)
		}

	}

}

```




