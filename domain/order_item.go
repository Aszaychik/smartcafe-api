package domain

type OrderItem struct {
	ID        uint  `gorm:"int;primaryKey" json:"id"`
	OrderID   uint  `gorm:"int" json:"order_id"`
	ItemID    uint  `gorm:"int" json:"item_id"`
	Quantity  int   `gorm:"int" json:"quantity"`
	ItemPrice int   `gorm:"type:decimal(10,2)"`
	Order     Order `gorm:"foreignKey:OrderID" json:"order"`
	Item      Menu  `gorm:"foreignKey:ItemID" json:"item"`
}