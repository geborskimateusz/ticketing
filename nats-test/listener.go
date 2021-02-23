package main

import (
	"fmt"
	"log"
	"time"

	stan "github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("areyouthere", "124", stan.NatsURL("http://localhost:4222"))
	if err != nil {
		log.Fatal(err)
	}

	aw, _ := time.ParseDuration("60s")
	for {
		_, err = sc.Subscribe("areyouthere:chat-invitation", func(m *stan.Msg) {
			m.Ack()
			fmt.Printf("Received a message: %s\n", string(m.Data))
		}, stan.SetManualAckMode(), stan.AckWait(aw))
	}

}
