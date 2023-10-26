package domain

type Menu struct {
	ID              uint     `gorm:"type:int;primarykey" json:"id"`
	ItemName        string   `gorm:"type:varchar(255);uniqueIndex" json:"item_name"`
	ItemPrice       int      `gorm:"type:int" json:"item_price"`
	ItemDescription string   `gorm:"type:varchar(255)" json:"item_description"`
	ItemImage       string   `gorm:"type:varchar(255)" json:"item_image"`
	CategoryID      uint     `gorm:"type:int" json:"category_id"`
	Category        Category `gorm:"foreignKey:CategoryID" json:"category"`
}