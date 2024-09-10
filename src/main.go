package main

import (
	"fmt"

	dbconnection "github.com/mohammadrahimi/Inventory_Service_RabbitMQ_Go/src/Infrastructure/Persistence.Sql/DbConnection"
	repository "github.com/mohammadrahimi/Inventory_Service_RabbitMQ_Go/src/Infrastructure/Persistence.Sql/Repository/Stock"
	rabbitMQ "github.com/mohammadrahimi/Inventory_Service_RabbitMQ_Go/src/RabbitMQ"
)

var (
	repo repository.IStockRepository
)


func init() {

	DbConnection := dbconnection.NewSQLConnection("sqlserver://@192.168.56.1?database=Inventory")
	DB, err := DbConnection.DBSQL()
	if err != nil {
		panic(" ConnectionSql is Error !  " + DB.Name() + err.Error())
	}
	repo = repository.NewStockRepository(DB)

}

 
func main() {
 
	
	 fmt.Println(" Consumer RabiitMQ  Start .... " )
	 rabbitMQ.Consume(repo,"inventoryOrder")
	 
}

  