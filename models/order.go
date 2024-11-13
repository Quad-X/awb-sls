package models

type Address struct {
	Validated struct {
		Type       string  `json:"string"`
		PhoneNo    string  `json:"phone_number"`
		MobileNo   string  `json:"mobile_number"`
		Line1      string  `json:"line_1"`
		Line2      *string `json:"line_2"`
		District   string  `json:"district"`
		City       string  `json:"city"`
		PostalCode string  `json:"postal_code"`
		Remarks    string  `json:"remarks"`
		Country    string  `json:"country"`
		HubCode    string  `json:"hub_code"`
		HubName    string  `json:"hub_alias"`
	} `json:"v_address"`
}

type Order struct {
	TrackingNumber  string   `json:"tracking_number"`
	PickupAddress   Address  `json:"pickup_address"`
	DeliveryAddress Address  `json:"delivery_address"`
	ReturnAddress   *Address `json:"return_address"`
}
