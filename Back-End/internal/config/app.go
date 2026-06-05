package config

import (
	"lotcastick-backend/internal/delivery/http/route"
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
	routeConfig := route.RouteConfig{
		Router: bsConfig.Router,
	}
	routeConfig.SetupRoutes()
}
