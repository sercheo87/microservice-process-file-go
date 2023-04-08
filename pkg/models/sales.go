package models

import (
	"gorm.io/gorm"
	"time"
)

type Sales struct {
	Region        string  `json:"region"`
	Country       string  `json:"country"`
	ItemType      string  `json:"item_type"`
	SaleChannel   string  `json:"sale_channel"`
	OrderPriority string  `json:"order_priority"`
	OrderDate     string  `json:"order_date"`
	OrderId       int64   `json:"order_id"`
	ShipDate      string  `json:"ship_date"`
	UnitSold      int64   `json:"unit_sold"`
	UnitPrice     float64 `json:"unit_price"`
	UnitCost      float64 `json:"unit_cost"`
	TotalRevenue  float64 `json:"total_revenue"`
	TotalCost     float64 `json:"total_cost"`
	TotalProfit   float64 `json:"total_profit"`
}

type SalesModel struct {
	gorm.Model
	OrderID   int64          `gorm:"primaryKey"`
	Region    string         `gorm:"column:region"`
	UnitSold  int64          `gorm:"column:unit_sold"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (SalesModel) TableName() string {
	return "sales"
}
