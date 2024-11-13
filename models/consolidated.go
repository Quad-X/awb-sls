package models

type Hub struct {
	Name      string `json:"name"`
	Line1     string `json:"line_1"`
	Line2     string `json:"line_2"`
	District  string `json:"district"`
	City      string `json:"city"`
	ContactNo string `json:"contact_no"`
}

type Consolidated struct {
	TrackingNumber string `json:"tracking_number"`
	SenderName     string `json:"sender_name"`
	SenderHub      Hub    `json:"sender_hub"`
	ReceiverName   string `json:"receiver_name"`
	ReceiverHub    Hub    `json:"receiver_hub"`
	QRCode         *string
	BarCode        *string
}
