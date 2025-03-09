package server

import (
	"encoding/json"
	myapp "github.com/tubfuzzy/banraiphisan-reservation/internal/app"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	fiberLog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/tubfuzzy/banraiphisan-reservation/config"

	cachePkg "github.com/tubfuzzy/banraiphisan-reservation/pkg/cache"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/common/exception"
	dbPkg "github.com/tubfuzzy/banraiphisan-reservation/pkg/db"
	loggerPkg "github.com/tubfuzzy/banraiphisan-reservation/pkg/logger"
	minioPkg "github.com/tubfuzzy/banraiphisan-reservation/pkg/minio"
)

type Server struct {
	app    *fiber.App
	conf   *config.Configuration
	db     *dbPkg.DB
	cache  cachePkg.Engine
	logger loggerPkg.Logger
	minio  *minioPkg.MinioClient
}

func New() (*Server, error) {
	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	logger := loggerPkg.NewLogger(conf)
	cacheEngine, err := cachePkg.NewCache(conf)
	if err != nil {
		return nil, err
	}
	db, err := dbPkg.NewDB(conf.Database)
	if err != nil {
		return nil, err
	}
	minioClient, err := minioPkg.New(conf.Minio)
	if err != nil {
		return nil, err
	}
	app := NewFiberApp(conf, logger, cacheEngine, db, minioClient)

	return &Server{
		app:    app,
		conf:   conf,
		db:     db,
		cache:  cacheEngine,
		logger: logger,
		minio:  minioClient,
	}, nil
}

func NewFiberApp(
	conf *config.Configuration,
	logger loggerPkg.Logger,
	cacheEngine cachePkg.Engine,
	db *dbPkg.DB,
	minioClient *minioPkg.MinioClient,
) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
		ReadTimeout:  time.Second * conf.Server.ReadTimeout,
		WriteTimeout: time.Second * conf.Server.WriteTimeout,
		JSONDecoder:  json.Unmarshal,
		JSONEncoder:  json.Marshal,
	})

	app.Use(cors.New())
	app.Use(etag.New())
	app.Use(recover.New())

	app.Use(fiberLog.New(fiberLog.Config{
		Next:         nil,
		Done:         nil,
		Format:       "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       os.Stdout,
	}))

	basePath := conf.Server.BaseURI
	api := app.Group(basePath + "/api")
	myapp.NewApplication(api, logger, db, cacheEngine, conf, minioClient)

	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "ok",
			"message": "Health check successful",
		})
	})

	app.Use(func(c *fiber.Ctx) error {
		panic(exception.NotFoundError{Message: "path " + c.Path() + " does not exist."})
	})

	return app
}

func (serv Server) App() *fiber.App {
	return serv.app
}

func (serv Server) Config() *config.Configuration {
	return serv.conf
}

func (serv Server) Logger() loggerPkg.Logger {
	return serv.logger
}

func (serv Server) DB() *dbPkg.DB {
	return serv.db
}

func (serv Server) Cache() cachePkg.Engine {
	return serv.cache
}

func (serv Server) Minio() *minioPkg.MinioClient {
	return serv.minio
}
