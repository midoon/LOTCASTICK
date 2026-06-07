package route

import (
	"lotcastick-backend/internal/controller"
	"net/http"

	"github.com/gorilla/mux"
)

type RouteConfig struct {
	Router         *mux.Router
	UserController *controller.UserController
}

func (rc *RouteConfig) SetupRoutes() {
	api := rc.Router.PathPrefix("/api").Subrouter()

	rc.setupPublicRoutes(api)

	// private := api.PathPrefix("").Subrouter()
	// private.Use(AuthMiddleware)

	// rc.setupPrivateRoutes(private)
}

func (rc *RouteConfig) setupPublicRoutes(api *mux.Router) {
	api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Health route"))
	}).Methods("GET")

	api.HandleFunc("/register", rc.UserController.Register).Methods("POST")
	api.HandleFunc("/login", rc.UserController.Login).Methods("POST")
}

func (rc *RouteConfig) setupPrivateRoutes(private *mux.Router) {

}
