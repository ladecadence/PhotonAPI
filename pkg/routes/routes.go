package routes

import (
	"net/http"

	"github.com/ladecadence/PhotonAPI/pkg/config"
	"github.com/ladecadence/PhotonAPI/pkg/controllers"
	"github.com/ladecadence/PhotonAPI/pkg/database"
)

func RegisterRoutes(db database.Database, config config.Config, router *http.ServeMux) {
	router.HandleFunc("GET /api", controllers.ConfMiddleWare(db, config, controllers.ApiRoot))
	router.HandleFunc("GET /api/walls", controllers.ConfMiddleWare(db, config, controllers.ApiGetWalls))
	router.HandleFunc("GET /api/wall/{uid}", controllers.ConfMiddleWare(db, config, controllers.ApiGetWall))
	router.HandleFunc("POST /api/newwall", controllers.ConfMiddleWare(db, config, controllers.ApiNewWall))
}
