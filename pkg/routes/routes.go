package routes

import (
	"net/http"

	"github.com/ladecadence/PhotonAPI/pkg/config"
	"github.com/ladecadence/PhotonAPI/pkg/controllers"
	"github.com/ladecadence/PhotonAPI/pkg/database"
)

func RegisterRoutes(db database.Database, config config.Config, router *http.ServeMux) {
	router.HandleFunc("GET /api", controllers.ConfMiddleWare(db, config, controllers.ApiRoot))
	// walls
	router.HandleFunc("GET /api/walls", controllers.ConfMiddleWare(db, config, controllers.ApiGetWalls))
	router.HandleFunc("GET /api/wall/{uid}", controllers.ConfMiddleWare(db, config, controllers.ApiGetWall))
	router.HandleFunc("POST /api/newwall", controllers.ConfMiddleWare(db, config, controllers.ApiNewWall))
	// problems
	router.HandleFunc("GET /api/problems", controllers.ConfMiddleWare(db, config, controllers.ApiGetProblems))
	router.HandleFunc("GET /api/problem/{uid}", controllers.ConfMiddleWare(db, config, controllers.ApiGetProblem))
	router.HandleFunc("POST /api/newproblem", controllers.ConfMiddleWare(db, config, controllers.ApiNewProblem))
	router.HandleFunc("GET /api/wall/{walluid}/problems", controllers.ConfMiddleWare(db, config, controllers.ApiGetWallProblems))
	// users
	router.HandleFunc("POST /api/signup", controllers.ConfMiddleWare(db, config, controllers.ApiSignup))
	router.HandleFunc("GET /api/login", controllers.ConfMiddleWare(db, config, controllers.ApiLogin))
	router.HandleFunc("GET /api/logout", controllers.ConfMiddleWare(db, config, controllers.ApiLogout))
}
