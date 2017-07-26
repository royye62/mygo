package main

import (
    "fmt"
    "time"

    "github.com/samuel/go-zookeeper/zk"
)

func must(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    conn, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
    must(err)
    defer conn.Close()

    //children, stat, ch, err := conn.ChildrenW("/")
    children, stat, _, err := conn.ChildrenW("/")
    must(err)
    fmt.Printf("%+v %+v\n", children, stat)

    flags := int32(0)
    acl := zk.WorldACL(zk.PermAll)

    exists, stat, err := conn.Exists("/01")
    must(err)
    fmt.Printf("exists: %+v %+v\n", exists, stat)

    if !exists {
        path, err := conn.Create("/01", []byte("data"), flags, acl)
        must(err)
        fmt.Printf("create: %+v\n", path)
    }

    data, stat, err := conn.Get("/01")
    must(err)
    fmt.Printf("get:    %+v %+v\n", string(data), stat)

    stat, err = conn.Set("/01", []byte("newdata"), stat.Version)
    must(err)
    fmt.Printf("set:    %+v\n", stat)

//   e := <-ch
//   fmt.Printf("%+v\n", e)
    ch2 := make(chan int)
    <-ch2
}
