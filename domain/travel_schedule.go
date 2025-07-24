package domain

type TravelSchRequest struct {
	ID          int    `gorm:"column:id" json:"id"`
	UserID      int    `gorm:"column:user_id" json:"user_id"`
	Location    int    `gorm:"column:locations" json:"location"`
	PeriodStart string `gorm:"column:period_start" json:"period_start"`
	PeriodEnd   string `gorm:"column:period_end" json:"period_end"`
	Status      int    `gorm:"column:status" json:"status"`
}

type TravelSchResponse struct {
	ID          int    `gorm:"column:id" json:"id"`
	Location    string `gorm:"column:locations" json:"location"`
	PeriodStart string `gorm:"column:period_start" json:"period_start"`
	PeriodEnd   string `gorm:"column:period_end" json:"period_end"`
	Status      int    `gorm:"status" json:"status"`
}
