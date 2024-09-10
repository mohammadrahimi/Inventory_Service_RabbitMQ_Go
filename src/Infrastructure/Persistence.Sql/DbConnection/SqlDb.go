package dbconnection

import (
	 

	 models "github.com/mohammadrahimi/Inventory_Service_RabbitMQ_Go/src/Infrastructure/Persistence.Sql/Models"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type DataBaseSql interface {
	DBSQL() *gorm.DB
}

type SQLConnection struct {
	ConnectionString string
}

func NewSQLConnection(ConnectionString string) *SQLConnection{
   return &SQLConnection{
	     ConnectionString: ConnectionString,
   }
}

func(cn *SQLConnection) DBSQL() (*gorm.DB,error){

    var db *gorm.DB 

	db, err := gorm.Open(sqlserver.Open(cn.ConnectionString), &gorm.Config{})
    if err != nil { 
      panic("failed to connect database") 
    } 

   err = db.Table("Stock").AutoMigrate(&models.StockEntity{}) 
   if err != nil { 
     return  nil,err
    } 
      
    return db,nil
}