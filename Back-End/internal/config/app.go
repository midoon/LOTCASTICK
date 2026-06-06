package config

import (
	"lotcastick-backend/internal/controller"
	"lotcastick-backend/internal/delivery/http/route"
	"lotcastick-backend/internal/repository"
	"lotcastick-backend/internal/usecase"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	ViperConfig *viper.Viper
	Database    *gorm.DB
	Router      *mux.Router
	HttpClient  *http.Client
	Validate    *validator.Validate
	RedisClient *redis.Client
}

func NewBootstrapConfig(bsConfig *BootstrapConfig) {

	userRepo := repository.NewUserRepository(bsConfig.Database)
	tokenRepo := repository.NewRefreshTokenRepository(bsConfig.Database)

	userUsecase := usecase.NewUserUsecase(userRepo, tokenRepo, bsConfig.Validate, bsConfig.ViperConfig)
	userController := controller.NewUserController(userUsecase)

	routeConfig := route.RouteConfig{
		Router:         bsConfig.Router,
		UserController: userController,
	}
	routeConfig.SetupRoutes()
}
