package rabbitMQ

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	repository "github.com/mohammadrahimi/Inventory_Service_RabbitMQ_Go/src/Infrastructure/Persistence.Sql/Repository/Stock"
)

var (
	_repository repository.IStockRepository
)

func Consume(repository  repository.IStockRepository,QueueName string) {

	 _repository =repository

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

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	FailOnError(err, "Failed to declare a queue")

	stopChan := make(chan bool)

	go func() {
		for d := range msgs {

			reciveBody(d.Body)
			// d.Ack(false)
		}
	}()


	<-stopChan
	
}

 

func reciveBody(BodyRescive []byte) {

	var data map[string]interface{}
	if err := json.Unmarshal(BodyRescive, &data); err != nil {
		panic(err)
	}

	customerId := data["CustomerId"].(string)
	orderId :=data["OrderId"].(string)

	orderItems := data["OrderItems"].([]interface{})

	items := make([]string, len(orderItems))
	 
	products := make(map[string]int)

	for i, value := range orderItems {

		items[i] = fmt.Sprint(value)
		ss := strings.Split(items[i], " ")

        var pid string
		 
		for i, v := range ss {
				if i == 2 { // productId 
					values := strings.Split(v, ":")
					if values[1] != "" {
						pid = values[1]
					}
				}
				if i==3 { // quantity
					values := strings.Split(v, ":")
					valueQ := strings.Split(values[1], "]")
					quantity, err := strconv.Atoi(valueQ[0])
					if err != nil {
						fmt.Println("Not Convert string to int" , err)
					}

					products[pid] =  quantity
                    pid=""
				}
				
			}
		}
	 
		 

	var stockState bool = true

	for id, quantity := range products {
		stock, err := _repository.FindByProductId(id)
		if err != nil {
			fmt.Println(" Not Found Stock for = ", err)
		}
	  
		if stock.Quantity < quantity {
			stockState = false
			break
		}
	}

	
    bodySend := "{\"CustomerId\":\""+customerId+"\",\"OrderId\":\""+orderId+"\"}"

	if stockState {
		 Puplish(bodySend,"InventoryOrderApproved") 
	} else {
		Puplish(bodySend,"InventoryOrderRejected") 
	}

}
