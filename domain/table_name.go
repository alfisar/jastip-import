package domain

func (AddressOrder) TableName() string {
	return "address_order"
}

func (TravelSchResponse) TableName() string {
	return "traveler_schedule"
}

func (OrderDetail) TableName() string {
	return "orders_detail"
}
