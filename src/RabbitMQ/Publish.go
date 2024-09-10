package rabbitMQ

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

 

func  Puplish(Body string,QueueName string) {

	conn, ch := ConnectMQ()
	defer CloseMQ(conn, ch)

	 
	q, err := ch.QueueDeclare(
		QueueName,  
		false,        
		false,        
		false,        
		false,       
		nil,          
	)

	FailOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := Body
	err = ch.PublishWithContext(ctx,
		"",      
		q.Name,  
		false,   
		false,   
		amqp.Publishing{
			ContentType: "application/json", //"text/plain"
			Body:        []byte(body),
		})

	FailOnError(err, "Failed to publish a message")
	 
}