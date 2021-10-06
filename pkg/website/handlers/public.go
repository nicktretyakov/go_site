package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	
)

type PublicHandler struct {
	logger   *logrus.Logger
	videRepo repository.VideoRepo
}

// NewPublicHandler initialize new handler
func NewPublicHandler(videoRepo repository.VideoRepo, logger *logrus.Logger) *PublicHandler {
	return &PublicHandler{}
}

// PopularAuthors return list of popular authors
func (p *PublicHandler) PopularAuthors(c *fiber.Ctx) (interface{}, error) {
	return "ok", nil
}
