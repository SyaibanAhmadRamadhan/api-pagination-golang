package httpprotocol

import (
	"fmt"

	"github.com/SyaibanAhmadRamadhan/technical-test-pt-zahir-international/config"
	httphand "github.com/SyaibanAhmadRamadhan/technical-test-pt-zahir-international/src/handlers/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rs/zerolog/log"
)

type HttpImpl struct {
	HttpHandler *httphand.HttpHandlerImpl
}

func NewHttpImpl(
	httpHandler *httphand.HttpHandlerImpl,
) *HttpImpl {
	return &HttpImpl{
		HttpHandler: httpHandler,
	}
}

func (h *HttpImpl) Listen() {
	app := fiber.New()

	app.Use(
		logger.New(),
	)

	h.HttpHandler.HttpHandlerRouter(app)

	port := fmt.Sprintf(":%s", config.Get().Application.Port)
	if err := app.Listen(port); err != nil {
		panic(err)
	}

	log.Info().Msgf("server started on port %s", port)
}
