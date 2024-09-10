package models

import (
	"time"

	"github.com/google/uuid"
)

type StockEntity struct {
	Id            uuid.UUID  `gorm:"primaryKey"` // ";unique"
	Amount        float64
	Currency      string     `gorm:"type:nvarchar(50);NOT NULL"`
	Quantity      int
	CreatedAt     time.Time
	ProductId     uuid.UUID
}