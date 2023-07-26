package main
import (
    "fmt"
    "sync"
)
var wg sync.WaitGroup

func main() {
    c1 := make(chan string)
    wg.Add(2)
    go func() {
        defer wg.Done()
        c1 <- "one"
    }()
    go func() {
        defer wg.Done()
        fmt.Println("message"+<-c1)
    }()
    wg.Wait()
}
