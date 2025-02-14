package domain

type TravelSchRequest struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Location    string `json:"location"`
	PeriodStart string `json:"period_start"`
	PeriodEnd   string `json:"period_end"`
	Status      int    `json:"status"`
}

type TravelSchResponse struct {
	ID          int    `json:"id"`
	Location    string `json:"location"`
	PeriodStart string `json:"period_start"`
	PeriodEnd   string `json:"period_end"`
}
