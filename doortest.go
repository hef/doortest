package main

import (
    "fmt"
    "time"
    zmq "github.com/pebbe/zmq4"
)

func main() {
    publisher, err := zmq.NewSocket(zmq.PUB)
    if err != nil {
        fmt.Println(err)
    }

    publisher.Bind("tcp://*:5556")
    time.Sleep(time.Second)
    for i := 0; i < 1000; i++ {
        time.Sleep(time.Second)
        _, err := publisher.SendMessage("door.test", fmt.Sprintf("%d-%s", i, "Open"))
        if err != nil {
            fmt.Println(err)
        }
    }
}
