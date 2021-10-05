package httputils

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

// FiberJSONWrapper wrapper for response json
func FiberJSONWrapper(callback func(c *fiber.Ctx) (interface{}, error)) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		res, err := callback(c)
		if err != nil {
			return err
		}

		return json.NewEncoder(c.Type("json", "utf-8").Response().BodyWriter()).Encode(res)
	}
}
