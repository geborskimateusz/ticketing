package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	stan "github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("areyouthere", "124", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatal(err)
	}

	aw, _ := time.ParseDuration("60s")

	messageNum := 0
	_, err = sc.Subscribe("areyouthere:chat", func(m *stan.Msg) {
		m.Ack()
		messageNum++
		fmt.Printf("Received a message %d: %s\n", messageNum, string(m.Data))
	}, stan.SetManualAckMode(), stan.AckWait(aw))

	runtime.Goexit()
}
