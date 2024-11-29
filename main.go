package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to the first NATS server (30041)
	nc1, err := nats.Connect("nats://localhost:30041")
	if err != nil {
		log.Fatalf("Error connecting to NATS server on port 30041: %v", err)
	}
	defer nc1.Close()

	// Connect to the second NATS server (30042)
	nc2, err := nats.Connect("nats://localhost:30042")
	if err != nil {
		log.Fatalf("Error connecting to NATS server on port 30042: %v", err)
	}
	defer nc2.Close()

	// Subscribe to all subjects using wildcard ">"
	_, err = nc1.Subscribe(">", func(msg *nats.Msg) {
		// Print the received message for debugging
		fmt.Printf("Received message on subject %s: %s\n", msg.Subject, string(msg.Data))

		// Publish the received message to the second NATS server
		err := nc2.Publish(msg.Subject, msg.Data)
		if err != nil {
			log.Printf("Error publishing message to second server: %v", err)
		} else {
			fmt.Printf("Message replicated to %s on second NATS server\n", msg.Subject)
		}
	})
	if err != nil {
		log.Fatalf("Error subscribing to all subjects: %v", err)
	}

	// Keep the program running to continue listening for events
	fmt.Println("Listening for all events... Press CTRL+C to exit.")
	select {}
}
