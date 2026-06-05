package main

import (
	"fmt"
	"lotcastick-backend/internal/config"
	"lotcastick-backend/internal/delivery/http/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	validate := config.NewValidator()

	r := mux.NewRouter()
	httpClient := http.Client{}

	bsConfig := &config.BootstrapConfig{
		ViperConfig: viperConfig,
		Database:    db,
		Router:      r,
		HttpClient:  &httpClient,
		Validate:    validate,
	}

	config.NewBootstrapConfig(bsConfig)

	handler := middleware.CorsMiddleware(r)

	addr := fmt.Sprintf("%s:%s", viperConfig.GetString("web.host"), viperConfig.GetString("web.port"))
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	fmt.Println("Server is running on port", addr)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
