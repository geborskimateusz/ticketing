package main

import stan "github.com/nats-io/stan.go"

func main() {
	sc, err := stan.Connect("aretouthere", "abc", stan.NatsURL("http://localhost:4222"))

}
