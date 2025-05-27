package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/banraiphisan/banraicore-api/config"
	"github.com/banraiphisan/banraicore-api/pkg/cache"
	"github.com/banraiphisan/banraicore-api/pkg/db"
	"github.com/banraiphisan/banraicore-api/pkg/logger"
	minioPkg "github.com/banraiphisan/banraicore-api/pkg/minio"

	authhandler "github.com/banraiphisan/banraicore-api/internal/usecase/auth/controller/http"
	authrepository "github.com/banraiphisan/banraicore-api/internal/usecase/auth/repository"
	authservice "github.com/banraiphisan/banraicore-api/internal/usecase/auth/service"

	reporthandler "github.com/banraiphisan/banraicore-api/internal/usecase/report/controller/http"
	reportrepository "github.com/banraiphisan/banraicore-api/internal/usecase/report/repository"
	reportservice "github.com/banraiphisan/banraicore-api/internal/usecase/report/service"

	reservationhandler "github.com/banraiphisan/banraicore-api/internal/usecase/reservation/controller/http"
	reservationrepository "github.com/banraiphisan/banraicore-api/internal/usecase/reservation/repository"
	reservationservice "github.com/banraiphisan/banraicore-api/internal/usecase/reservation/service"

	roomhandler "github.com/banraiphisan/banraicore-api/internal/usecase/room/controller/http"
	roomrepository "github.com/banraiphisan/banraicore-api/internal/usecase/room/repository"
	roomservice "github.com/banraiphisan/banraicore-api/internal/usecase/room/service"

	userhandler "github.com/banraiphisan/banraicore-api/internal/usecase/user/controller/http"
	userrepository "github.com/banraiphisan/banraicore-api/internal/usecase/user/repository"
	userservice "github.com/banraiphisan/banraicore-api/internal/usecase/user/service"

	shorturlhandler "github.com/banraiphisan/banraicore-api/internal/usecase/shorturl/controller/http"
	shorturlrepository "github.com/banraiphisan/banraicore-api/internal/usecase/shorturl/repository"
	shorturlservice "github.com/banraiphisan/banraicore-api/internal/usecase/shorturl/service"
)

func NewApplication(api fiber.Router, logger logger.Logger, db *db.DB, cache cache.Engine, config *config.Configuration, minioClient *minioPkg.MinioClient) {
	v1 := api.Group("/v1")

	authRepository := authrepository.NewAuthRepository(db, logger, cache, config)
	authService := authservice.NewAuthService(authRepository, cache, logger, config)
	authHandler := authhandler.NewAuthHandler(authService, config)
	authHandler.InitRoute(v1)

	reportRepository := reportrepository.NewReportRepository(db, logger, cache, config)
	reportService := reportservice.NewReportService(reportRepository, cache, logger, config)
	reportHandler := reporthandler.NewReportHandler(reportService, config)
	reportHandler.InitRoute(v1)

	reservationRepository := reservationrepository.NewReservationRepository(db, logger, cache, config)
	reservationService := reservationservice.NewReservationService(reservationRepository, cache, logger, config)
	reservationHandler := reservationhandler.NewReservationHandler(reservationService, config)
	reservationHandler.InitRoute(v1)

	roomRepository := roomrepository.NewRoomRepository(db, logger, cache, config)
	roomService := roomservice.NewRoomService(roomRepository, cache, logger, config)
	roomHandler := roomhandler.NewRoomHandler(roomService, config)
	roomHandler.InitRoute(v1)

	userRepository := userrepository.NewUserRepository(db, logger, cache, config)
	userService := userservice.NewUserService(userRepository, cache, logger, config)
	userHandler := userhandler.NewUserHandler(userService, config)
	userHandler.InitRoute(v1)

	shortUrlRepository := shorturlrepository.NewShortUrlRepository(db, logger, cache, config)
	shortUrlService := shorturlservice.NewShortUrlService(shortUrlRepository, cache, logger, config)
	shortUrlHandler := shorturlhandler.NewShortUrlHandler(shortUrlService, config)
	shortUrlHandler.InitRoute(v1)

}
