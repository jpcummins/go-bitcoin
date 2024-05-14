package main

import (
	"log"

	"github.com/jpcummins/go-bitcoin"
)

func main() {
	zmq := bitcoin.NewZMQ("localhost", 28332)

	ch := make(chan []string)

	go func() {
		for c := range ch {
			log.Printf("%v", c)
		}
	}()

	done := make(chan bool)
	if err := zmq.Subscribe("hashtx", ch, done); err != nil {
		log.Fatalf("%v", err)
	}

	if err := zmq.Subscribe("hashblock", ch, done); err != nil {
		log.Fatalf("%v", err)
	}

	waitCh := make(chan bool)
	<-waitCh
}
