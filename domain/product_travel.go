package domain

type ProductsTravel struct {
	ID        int `json:"column:id" gorm:"id"`
	ProductID int `json:"column:product_id" gorm:"product_id"`
	TravelID  int `json:"column:treveler_schedule_id" gorm:"treveler_schedule_id"`
}
