package domain

type OrderRequest struct {
	ID            int           `json:"id" gorm:"id"`
	BuyerID       int           `json:"buyer_id" gorm:"buyer_id"`
	TravelID      int           `json:"travel_schedule_id" gorm:"travel_schedule_id"`
	AddressID     int           `json:"address_id" gorm:"address_id"`
	Price         float32       `json:"price" gorm:"price"`
	Status        int           `json:"status" gorm:"status"`
	PaymentStatus int           `json:"payment_satatus" gorm:"payment_satatus"`
	PaymentMethod int           `json:"payment_status" gorm:"payment_status"`
	Product       []ProductData `json:"products"`
}

type OrderOneResponse struct {
	ID            int               `json:"id" gorm:"primaryKey; id"`
	Invoice       string            `json:"invoice" gorm:"invoice"`
	AddressID     int               `json:"-" gorm:"address_id"`
	TravelID      int               `json:"-" gorm:"travel_schedule_id"`
	Price         float32           `json:"price" gorm:"price"`
	Status        int               `json:"status" gorm:"status"`
	PaymentStatus int               `json:"payment_satatus" gorm:"payment_satatus"`
	PaymentMethod int               `json:"payment_status" gorm:"payment_status"`
	Travel        TravelSchResponse `json:"travel_schedule" gorm:"foreignKey:TravelID;references:ID"`
	Address       AddressOrder      `json:"address" gorm:"foreignKey:AddressID;references:ID"`
	Product       []OrderDetail     `json:"products" gorm:"foreignKey:OrderID;references:ID"`
}

type OrderListResponse struct {
	ID      int               `json:"id" gorm:"id"`
	Invoice string            `json:"invoice" gorm:"invoice"`
	Price   float32           `json:"price" gorm:"price"`
	Status  int               `json:"status" gorm:"status"`
	Travel  TravelSchResponse `json:"travel_schedule" gorm:"travel_schedule"`
	Address AddressOrder      `json:"address" gorm:"address"`
}

type OrderData struct {
	ID            int     `json:"id" gorm:"primaryKey; id"`
	BuyerID       int     `json:"buyer_id" gorm:"buyer_id"`
	TravelID      int     `json:"travel_schedule_id" gorm:"travel_schedule_id"`
	Invoice       string  `json:"invoice" gorm:"invoice"`
	AddressID     int     `json:"address_id" gorm:"address_id"`
	Price         float32 `json:"price" gorm:"price"`
	Status        int     `json:"status" gorm:"status"`
	PaymentStatus int     `json:"payment_satatus" gorm:"payment_satatus"`
	PaymentMethod int     `json:"payment_status" gorm:"payment_status"`
}

type OrderDetail struct {
	ID              int     `json:"id" gorm:"id"`
	OrderID         int     `json:"order_id" gorm:"order_id"`
	ProductID       int     `json:"products_id" gorm:"products_id"`
	ProductName     int     `json:"products_name" gorm:"products_name"`
	ProductImage    int     `json:"products_image" gorm:"products_image"`
	ProductPrice    float32 `json:"products_price" gorm:"products_price"`
	ProductQuantity int     `json:"products_quantity" gorm:"products_quantity"`
	Quantity        int     `json:"quantity" gorm:"quantity"`
	Price           float32 `json:"price" gorm:"price"`
}

type AddressOrder struct {
	Id            int    `gorm:"primaryKey; column:id" json:"id"`
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
