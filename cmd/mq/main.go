package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
)

func main() {
	// Get the connection string from the environment variable
	url := os.Getenv("AMQP_URL")

	//If it doesn't exist, use the default connection string.

	if url == "" {
		//Don't do this in production, this is for testing purposes only.
		url = "amqp://guest:guest@localhost:5672"
	}

	// Connect to the rabbitMQ instance
	connection, err := amqp.Dial(url)

	fmt.Printf("Conn here %+v \n", connection)



	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}
}
