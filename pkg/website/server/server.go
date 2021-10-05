package server

import (
	"github.com/gofiber/fiber/v2"

	"website.friendsonly.me/pkg/tools/httputils"
	"website.friendsonly.me/pkg/website/config"
	"website.friendsonly.me/pkg/website/handlers"
)

// BindHandlers provide
func BindHandlers(app *fiber.App, publicHandler *handlers.PublicHandler, cfg *config.Config) {

	public := app.Group("/public")
	public.Get("/popular_authors", httputils.FiberJSONWrapper(publicHandler.PopularAuthors))

}
