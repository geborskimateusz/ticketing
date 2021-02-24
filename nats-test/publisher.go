package main

import (
	"log"
	"time"

	stan "github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("areyouthere", "123",
		stan.NatsURL("nats://localhost:4222"),
		stan.MaxPubAcksInflight(1000))
	if err != nil {
		log.Fatal(err)
	}

	ackHandler := func(ackedNuid string, er error) {
		if er != nil {
			log.Printf("Warning: error publishing msg id %s: %v\n", ackedNuid, er.Error())
		} else {
			log.Printf("Received ack for msg id %s\n", ackedNuid)
		}
	}

	for {
		time.Sleep(5 * time.Second)
		_, err = sc.PublishAsync("areyouthere:chat", []byte("Hello World"), ackHandler)
		if err != nil {
			log.Fatal(err)
		}
	}

}
