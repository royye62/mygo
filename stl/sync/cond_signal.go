package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    locker := new(sync.Mutex)
    cond := sync.NewCond(locker)
    done := false

    cond.L.Lock()

    go func() {
        time.Sleep(2e9)
        done = true
        cond.Signal()
    }()

    if (!done) {
        cond.Wait()
    }

    fmt.Println("now done is ", done);
}
