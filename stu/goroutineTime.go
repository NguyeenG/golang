package main
import "fmt"
import "time"
import "sync"
func main(){
    var wg = sync.WaitGroup{}
    wg.Add(1)
    go func ()  {
        count("sleep")
        count("fish")
        wg.Done()
    }()
    wg.Wait()
    fmt.Println("cut")
}
func count (name string){
    for i:= 1;i<= 5;i++{
        fmt.Println(name,i)
        time.Sleep(1* time.Second)
    }
}
