package main

import (
    "fmt"
    "time"
    "encoding/json"
    zmq "github.com/pebbe/zmq4"
)

type EventMessage struct {

    EventType string
    Event string
    Time time.Time
}

func main() {
    publisher, err := zmq.NewSocket(zmq.PUB)
    if err != nil {
        fmt.Println(err)
    }

    publisher.Bind("tcp://*:5556")
    time.Sleep(time.Second)
    for i := 0; i < 1000; i++ {
        time.Sleep(time.Second)
        message, _ := json.Marshal(EventMessage{"door", "opened", time.Now()})
        _, err := publisher.SendBytes(message, 0)
        if err != nil {
            fmt.Println(err)
        }
    }
}
