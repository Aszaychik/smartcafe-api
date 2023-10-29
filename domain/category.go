package domain

type Category struct {
	ID                  uint   `gorm:"type:int;primarykey" json:"id"`
	CategoryName        string `gorm:"type:varchar(255)" json:"category_name"`
	CategoryDescription string `gorm:"type:text" json:"category_description"`
}