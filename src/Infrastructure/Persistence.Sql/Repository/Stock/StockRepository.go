package repository

import (
	 

	models "github.com/mohammadrahimi/Inventory_Service_RabbitMQ_Go/src/Infrastructure/Persistence.Sql/Models"
	"gorm.io/gorm"
)
 
type StockRepository struct {
	db    *gorm.DB
}

func NewStockRepository(db *gorm.DB)   IStockRepository{
	return &StockRepository{db: db}
}
 
func (s *StockRepository) FindById(id string) (models.StockEntity , error) {
	 
	var  stock   models.StockEntity 
	s.db.Table("Stock").Select("*").Where("id=?", id).Find(&stock)
    return  stock,nil  
	 
}

func (s *StockRepository) FindByProductId(id string) (models.StockEntity, error) {
	 
	var  stock   models.StockEntity 
	s.db.Table("Stock").Select("*").Where("product_id=?", id).Find(&stock)
     return  stock,nil  
	 
}
