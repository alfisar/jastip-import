package domain

type ProductData struct {
	ID       int     `gorm:"id" json:"id"`
	Name     string  `gorm:"name" json:"name"`
	Desc     string  `gorm:"desc" json:"desc"`
	Price    float64 `gorm:"price" json:"price"`
	TravelID int     `gorm:"travel_schedule_id" json:"travel_schedule_id"`
	Image    string  `gorm:"image" json:"image"`
	Status   int     `gorm:"status" json:"status"`
	Quantity int     `gorm:"quantity" json:"quantity"`
}

type ProductResp struct {
	ID                int     `gorm:"id" json:"id"`
	Name              string  `gorm:"name" json:"name"`
	Desc              string  `gorm:"desc" json:"desc"`
	Price             float64 `gorm:"price" json:"price"`
	TravelID          int     `gorm:"travel_schedule_id" json:"travel_schedule_id"`
	TravelLocation    int     `gorm:"location" json:"location"`
	TravelPeriodStart string  `gorm:"column:period_start" json:"period_start"`
	TravelPeriodEnd   string  `gorm:"column:period_end" json:"period_end"`
	Image             string  `gorm:"image" json:"image"`
	Status            int     `gorm:"status" json:"status"`
	Quantity          int     `gorm:"quantity" json:"quantity"`
}
