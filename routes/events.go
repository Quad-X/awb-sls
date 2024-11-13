package routes

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	lib "github.com/quad-x/awb-sls/lib"
	m "github.com/quad-x/awb-sls/models"
)

// Receive address validation events from AMS via SQS
func ReceiveEvents(ctx *fiber.Ctx) error {
	var event m.Event
	// parse sqs event
	if err := ctx.BodyParser(&event); err != nil {
		return err
	}
	// iterate records from sqs event
	for _, record := range event.Records {
		var order m.Order
		// parse order from address validation event
		if err := json.Unmarshal([]byte(record.Body), &order); err != nil {
			return err
		}
		// generate qr code
		lib.GenerateQR(order.TrackingNumber, 83, 83)
		// generate bar code
		lib.GenerateBarCode(order.TrackingNumber)
	}
	return ctx.JSON(event)
}
