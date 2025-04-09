package routes

import (
	"net/http"

	"github.com/ladecadence/PhotonAPI/pkg/config"
	"github.com/ladecadence/PhotonAPI/pkg/controllers"
	"github.com/ladecadence/PhotonAPI/pkg/database"
)

func RegisterRoutes(db database.Database, config config.Config, router *http.ServeMux) {
	router.HandleFunc("/api", controllers.ConfMiddleWare(db, config, controllers.ApiRoot))
	router.HandleFunc("/api/walls", controllers.ConfMiddleWare(db, config, controllers.ApiGetWalls))
}
