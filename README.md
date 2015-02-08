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
	watcher := gowatch.NewWatcher()

	watcher.Append("tcp", "www.google.com.tw:80", time.Second*5)
	watcher.Append("tcp", "localhost:3000", time.Second*1)

	//select {}
	for {
		select {
		    case node := <-watcher.WatchChan:
			    fmt.Println("watcher", node.s)
		}

	}

}
```




