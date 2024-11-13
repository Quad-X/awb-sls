package models

type Record struct {
	MessageId string `json:"messageId"`
	Body      string `json:"body"`
}

type Event struct {
	Records []Record `json:"records"`
}
