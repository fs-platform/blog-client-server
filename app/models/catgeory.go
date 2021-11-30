package models

type Category struct {
	BaseModel
	Name   string `gorm:"type:varchar(10);not null" json:"names"`
	Status uint8 `gorm:"type:uint;default:1;not null" json:"status"`
}
