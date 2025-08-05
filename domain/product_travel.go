package domain

type ProductsTravel struct {
	ID        int `json:"id" gorm:"column:id"`
	ProductID int `json:"product_id" gorm:"column:product_id"`
	TravelID  int `json:"treveler_schedule_id" gorm:"column:treveler_schedule_id"`
}

type ProductsTravelRequest struct {
	ID        int   `json:"id" gorm:"column:id"`
	ProductID []int `json:"product_id" gorm:"column:products_id"`
	TravelID  []int `json:"treveler_schedule_id" gorm:"column:treveler_schedules_id"`
}
