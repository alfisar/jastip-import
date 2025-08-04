package domain

type ProductData struct {
	ID       int     `gorm:"column:id" json:"id"`
	UserID   int     `gorm:"column:user_id" json:"user_id"`
	Name     string  `gorm:"column:name" json:"name"`
	Desc     string  `gorm:"column:desc" json:"desc"`
	Price    float64 `gorm:"column:price" json:"price"`
	Image    string  `gorm:"column:image" json:"image"`
	Status   int     `gorm:"column:status" json:"status"`
	Quantity int     `gorm:"column:quantity" json:"quantity"`
}

type ProductResp struct {
	ID       int     `gorm:"column:id" json:"id"`
	Name     string  `gorm:"column:name" json:"name"`
	Desc     string  `gorm:"column:desc" json:"desc"`
	Price    float64 `gorm:"column:price" json:"price"`
	Image    string  `gorm:"column:image" json:"image"`
	Status   int     `gorm:"column:status" json:"status"`
	Quantity int     `gorm:"column:quantity" json:"quantity"`
}
