package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tubfuzzy/banraiphisan-reservation/config"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/cache"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/db"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/logger"
)

func NewApplication(api fiber.Router, logger logger.Logger, db *db.DB, cache cache.Engine, config *config.Configuration) {

}
