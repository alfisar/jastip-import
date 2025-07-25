package domain

type ProductData struct {
	ID       int     `gorm:"id" json:"id"`
	UserID   int     `gorm:"user_id" json:"user_id"`
	Name     string  `gorm:"name" json:"name"`
	Desc     string  `gorm:"desc" json:"desc"`
	Price    float64 `gorm:"price" json:"price"`
	Image    string  `gorm:"image" json:"image"`
	Status   int     `gorm:"status" json:"status"`
	Quantity int     `gorm:"quantity" json:"quantity"`
}

type ProductResp struct {
	ID       int     `gorm:"id" json:"id"`
	Name     string  `gorm:"name" json:"name"`
	Desc     string  `gorm:"desc" json:"desc"`
	Price    float64 `gorm:"price" json:"price"`
	Image    string  `gorm:"image" json:"image"`
	Status   int     `gorm:"status" json:"status"`
	Quantity int     `gorm:"quantity" json:"quantity"`
}
