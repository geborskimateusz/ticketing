package main

import (
	"log"

	stan "github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("areyouthere", "123", stan.NatsURL("http://localhost:4222"))
	if err != nil {
		log.Fatal(err)
	}

	ackHandler := func(ackedNuid string, err error) {
		if err != nil {
			log.Printf("Warning: error publishing msg id %s: %v\n", ackedNuid, err.Error())
		} else {
			log.Printf("Received ack for msg id %s\n", ackedNuid)
		}
	}

	_, err = sc.PublishAsync("areyouthere:chat-invitation", []byte("Hello World"), ackHandler)
	if err != nil {
		log.Fatal(err)
	}

}
