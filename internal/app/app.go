package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tubfuzzy/banraiphisan-reservation/config"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/cache"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/db"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/logger"
	minioPkg "github.com/tubfuzzy/banraiphisan-reservation/pkg/minio"

	authhandler "github.com/tubfuzzy/banraiphisan-reservation/internal/usecase/auth/controller/http"
	authrepository "github.com/tubfuzzy/banraiphisan-reservation/internal/usecase/auth/repository"
	authservice "github.com/tubfuzzy/banraiphisan-reservation/internal/usecase/auth/service"

	reporthandler "github.com/tubfuzzy/banraiphisan-reservation/internal/usecase/report/controller/http"
	reportrepository "github.com/tubfuzzy/banraiphisan-reservation/internal/usecase/report/repository"
	reportservice "github.com/tubfuzzy/banraiphisan-reservation/internal/usecase/report/service"

	reservationhandler "github.com/tubfuzzy/banraiphisan-reservation/internal/usecase/reservation/controller/http"
	reservationrepository "github.com/tubfuzzy/banraiphisan-reservation/internal/usecase/reservation/repository"
	reservationservice "github.com/tubfuzzy/banraiphisan-reservation/internal/usecase/reservation/service"

	roomhandler "github.com/tubfuzzy/banraiphisan-reservation/internal/usecase/room/controller/http"
	roomrepository "github.com/tubfuzzy/banraiphisan-reservation/internal/usecase/room/repository"
	roomservice "github.com/tubfuzzy/banraiphisan-reservation/internal/usecase/room/service"

	userhandler "github.com/tubfuzzy/banraiphisan-reservation/internal/usecase/user/controller/http"
	userrepository "github.com/tubfuzzy/banraiphisan-reservation/internal/usecase/user/repository"
	userservice "github.com/tubfuzzy/banraiphisan-reservation/internal/usecase/user/service"

	shorturlhandler "github.com/tubfuzzy/banraiphisan-reservation/internal/usecase/shorturl/controller/http"
	shorturlrepository "github.com/tubfuzzy/banraiphisan-reservation/internal/usecase/shorturl/repository"
	shorturlservice "github.com/tubfuzzy/banraiphisan-reservation/internal/usecase/shorturl/service"
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
