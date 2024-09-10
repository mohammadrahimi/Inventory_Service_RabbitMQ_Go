package rabbitMQ

import (
	 
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectMQ() (*amqp.Connection, *amqp.Channel) {

	conn, err := amqp.Dial("amqp://eshop:6661@localhost:5672/orderVirtual")
	FailOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	// defer ch.Close()

	return conn, ch
}

func CloseMQ(conn *amqp.Connection, channel *amqp.Channel) {

	defer conn.Close()    //rabbit mq close
	defer channel.Close() //rabbit mq channel close
	
}