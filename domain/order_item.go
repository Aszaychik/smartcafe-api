package domain

type OrderItem struct {
	ID       uint `gorm:"int;primaryKey" json:"id"`
	OrderID  uint `gorm:"int" json:"order_id"`
	ItemID   uint `gorm:"int" json:"item_id"`
	Quantity int  `gorm:"int" json:"quantity"`
	Item     Menu `gorm:"foreignKey:ItemID" json:"item"`
}