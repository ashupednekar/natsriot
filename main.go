package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

func main() {

  sourceAddr := os.Getenv("SOURCE_SERVER_ADDR")
	nc1, err := nats.Connect(sourceAddr)
	if err != nil {
    log.Fatalf("Error connecting to source NATS server at %s -> %v", sourceAddr, err)
	}
	defer nc1.Close()

  destAddr := os.Getenv("DESTINATION_SERVER_ADDR")
	nc2, err := nats.Connect(destAddr)
	if err != nil {
    log.Fatalf("Error connecting to destination NATS server at %s:  %v", destAddr, err)
	}
	defer nc2.Close()

	// Subscribe to all subjects using wildcard ">"
	_, err = nc1.Subscribe(">", func(msg *nats.Msg) {
		// Print the received message for debugging
		fmt.Printf("Received message on subject %s: %s\n", msg.Subject, string(msg.Data))

		// Publish the received message to the second NATS server
		err := nc2.Publish(msg.Subject, msg.Data)
		if err != nil {
			log.Printf("Error publishing message to destination server: %v", err)
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
