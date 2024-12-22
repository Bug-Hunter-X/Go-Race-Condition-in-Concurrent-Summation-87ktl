```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var count int
        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        count += i
                }(i)
        }
        wg.Wait()
        fmt.Println(count) // This might print an unexpected result
}
```