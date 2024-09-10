package repository

import models "github.com/mohammadrahimi/Inventory_Service_RabbitMQ_Go/src/Infrastructure/Persistence.Sql/Models"

type IStockRepository interface {
	FindByProductId(id string) (models.StockEntity, error)
	FindById(id string) (models.StockEntity, error)
}