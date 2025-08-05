package domain

type ProductsTravel struct {
	ID        int `json:"id" gorm:"column:id"`
	ProductID int `json:"product_id" gorm:"column:product_id"`
	TravelID  int `json:"traveler_schedule_id" gorm:"column:traveler_schedule_id"`
}

type ProductsTravelRequest struct {
	ID        int   `json:"id" gorm:"column:id"`
	ProductID []int `json:"product_id" gorm:"column:products_id"`
	TravelID  []int `json:"traveler_schedule_id" gorm:"column:traveler_schedules_id"`
}
