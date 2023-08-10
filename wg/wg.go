
package main``
import (
    "fmt"
    "sync"
    "time"
)

var wg sync.WaitGroup


func worker(id int){
    defer wg.Done()
    fmt.Printf("worker %d starting \n",id)

    time.Sleep(time.Second)

    fmt.Printf("worker %d done \n",id)
}



func main(){
    
    size:=10
    wg.Add(size)

    for i:=0;i<10;i++{
        go worker(i)
    }

    wg.Wait()
    
    fmt.Printf("main Done\n")




}

