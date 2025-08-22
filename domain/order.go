package domain

type OrderRequest struct {
	ID            int           `json:"id" gorm:"column:id"`
	BuyerID       int           `json:"buyer_id" gorm:"column:buyer_id"`
	TravelID      int           `json:"travel_schedule_id" gorm:"column:travel_schedule_id"`
	AddressID     int           `json:"address_id" gorm:"column:address_id"`
	Price         float32       `json:"price" gorm:"column:price"`
	Status        int           `json:"status" gorm:"column:status"`
	PaymentStatus int           `json:"payment_status" gorm:"column:payment_status"`
	PaymentMethod int           `json:"payment_method" gorm:"column:payment_method"`
	Product       []ProductData `json:"products"`
}

type OrderOneResponse struct {
	ID            int               `json:"id" gorm:"column:primaryKey; id"`
	Invoice       string            `json:"invoice" gorm:"column:invoice"`
	AddressID     int               `json:"-" gorm:"column:address_id"`
	TravelID      int               `json:"-" gorm:"column:travel_schedule_id"`
	Price         float32           `json:"price" gorm:"column:price"`
	Status        int               `json:"status" gorm:"column:status"`
	PaymentStatus int               `json:"payment_satatus" gorm:"column:payment_satatus"`
	PaymentMethod int               `json:"payment_status" gorm:"column:payment_status"`
	Travel        TravelSchResponse `json:"travel_schedule" gorm:"foreignKey:TravelID;references:ID"`
	Address       AddressOrder      `json:"address" gorm:"foreignKey:AddressID;references:ID"`
	Product       []OrderDetail     `json:"products" gorm:"foreignKey:OrderID;references:ID"`
}

type OrderListResponse struct {
	ID                 int               `json:"id" gorm:"column:id"`
	Invoice            string            `json:"invoice" gorm:"column:invoice"`
	Price              float32           `json:"price" gorm:"column:price"`
	Status             int               `json:"status" gorm:"column:status"`
	AddressID          int               `json:"-" gorm:"column:address_id"`
	TravelerScheduleID int               `json:"-" gorm:"column:travel_schedule_id"`
	Travel             TravelSchResponse `json:"travel_schedule" gorm:"foreignKey:TravelerScheduleID;references:ID"`
	Address            AddressOrder      `json:"address" gorm:"foreignKey:AddressID;references:ID"`
}

type OrderData struct {
	ID            int     `json:"id" gorm:"primaryKey; id"`
	BuyerID       int     `json:"buyer_id" gorm:"column:buyer_id"`
	TravelID      int     `json:"travel_schedule_id" gorm:"column:traveler_schedule_id"`
	Invoice       string  `json:"invoice" gorm:"column:invoice"`
	AddressID     int     `json:"address_id" gorm:"column:address_id"`
	Price         float32 `json:"price" gorm:"column:price"`
	Status        int     `json:"status" gorm:"column:status"`
	PaymentStatus int     `json:"payment_status" gorm:"column:payment_status"`
	PaymentMethod int     `json:"payment_method" gorm:"column:payment_method"`
}

type OrderDetail struct {
	ID              int     `json:"id" gorm:"column:id"`
	OrderID         int     `json:"order_id" gorm:"column:order_id"`
	ProductID       int     `json:"products_id" gorm:"column:products_id"`
	ProductName     string  `json:"products_name" gorm:"column:products_name"`
	ProductImage    string  `json:"products_image" gorm:"column:products_image"`
	ProductPrice    float32 `json:"products_price" gorm:"column:products_price"`
	ProductQuantity int     `json:"products_quantity" gorm:"column:products_quantity"`
	Quantity        int     `json:"quantity" gorm:"column:quantity"`
	Price           float32 `json:"price" gorm:"column:price"`
}

type AddressOrder struct {
	ID            int    `gorm:"primaryKey; column:id" json:"id"`
	UserID        int    `gorm:"column:user_id" json:"user_id"`
	ReceiverName  string `gorm:"column:receiver_name" json:"receiver_name"`
	ReceiverPhone string `gorm:"column:receiver_phone" json:"receiver_phone"`
	Province      string `gorm:"column:province" json:"province"`
	Street        string `gorm:"column:street" json:"street"`
	City          string `gorm:"column:city" json:"city"`
	District      string `gorm:"column:district" json:"district"`
	SUbDistrict   string `gorm:"column:subdistrict" json:"subdistrict"`
	PostalCode    string `gorm:"column:postalcode" json:"postalcode"`
	Tag           string `gorm:"column:tag" json:"tag"`
}
