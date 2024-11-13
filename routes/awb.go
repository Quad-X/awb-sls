package routes

import (
	"bytes"

	"github.com/gofiber/fiber/v2"
	"github.com/quad-x/awb-sls/lib"
	m "github.com/quad-x/awb-sls/models"
)

// Receive address validation events from AMS
func GenerateConsolidatedAWB(ctx *fiber.Ctx) error {
	var err error
	var data m.Consolidated

	// parse payload
	if err = ctx.BodyParser(&data); err != nil {
		return err
	}

	var output []byte

	if output, err = lib.GenerateConsolidatedAWB(&data); err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	// Set Headers
	ctx.Set(fiber.HeaderContentType, "application/pdf")
	ctx.Set(fiber.HeaderContentDisposition, "attachment;")

	return ctx.SendStream(bytes.NewReader(output))
}
