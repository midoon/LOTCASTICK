package route

import (
	"lotcastick-backend/internal/controller"
	"lotcastick-backend/internal/delivery/http/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

type RouteConfig struct {
	Router               *mux.Router
	UserController       *controller.UserController
	JwtsecretKey         string
	SimulationController *controller.SimulationController
}

func (rc *RouteConfig) SetupRoutes() {
	api := rc.Router.PathPrefix("/api").Subrouter()

	rc.setupPublicRoutes(api)

	private := api.PathPrefix("").Subrouter()
	private.Use(func(next http.Handler) http.Handler {
		return middleware.AuthMiddleware(next, rc.JwtsecretKey)
	})

	rc.setupPrivateRoutes(private)
}

func (rc *RouteConfig) setupPublicRoutes(api *mux.Router) {
	api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Health route"))
	}).Methods("GET")

	api.HandleFunc("/register", rc.UserController.Register).Methods("POST")
	api.HandleFunc("/login", rc.UserController.Login).Methods("POST")
	api.HandleFunc("/refresh", rc.UserController.RefreshToken).Methods("POST")
}

func (rc *RouteConfig) setupPrivateRoutes(private *mux.Router) {
	private.HandleFunc("/logout", rc.UserController.Logout).Methods("POST")
	private.HandleFunc("/simulations", rc.SimulationController.CreateSimulation).Methods("POST")
}
