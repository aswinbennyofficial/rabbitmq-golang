package main

import (
  "log"

  amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
  if err != nil {
    log.Panicf("%s: %s", msg, err)
  }
}

func main() {
	// Establish connection to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	
	// Create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	
	// Declaring a queue
	q, err := ch.QueueDeclare(
	  "hello", // name
	  false,   // durable
	  false,   // delete when unused
	  false,   // exclusive
	  false,   // no-wait
	  nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Consume messages from the queue
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	  )
	  failOnError(err, "Failed to register a consumer")
	  
	  // Create a channel to receive messages from the queue
	  var forever chan struct{}
	  
	  // Create a goroutine to consume messages from the queue
	  go func() {
		for d := range msgs {
		  log.Printf("Received a message: %s", d.Body)
		}
	  }()
	  
	  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	 // Loop forever 
	  <-forever
}