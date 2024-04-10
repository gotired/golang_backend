package utils

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LogRequest(c *fiber.Ctx) error {
	start := time.Now()

	if err := c.Next(); err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		return err
	}

	resSize := c.Response().Header.ContentLength()

	responseTime := time.Since(start)

	log.Printf("[%s] - %s - \"%s %s\" %d %d %.2f ms",
		time.Now().Format("02/Jan/2006:15:04:05 -0700"), c.IP(), c.Method(), c.Path(),
		c.Response().StatusCode(), resSize, responseTime.Seconds()*1000)

	return nil
}
