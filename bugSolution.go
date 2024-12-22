```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var mu sync.Mutex // Add mutex for synchronization
        var count int
        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        mu.Lock() // Lock before accessing shared variable
                        count += i
                        mu.Unlock() // Unlock after accessing shared variable
                }(i)
        }
        wg.Wait()
        fmt.Println(count) // This will now print the correct sum
}
```