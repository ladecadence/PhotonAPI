package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ladecadence/PhotonAPI/pkg/config"
	"github.com/ladecadence/PhotonAPI/pkg/database"
)

var conf config.Config
var db database.Database

func ConfMiddleWare(dtb database.Database, c config.Config, h http.HandlerFunc) http.HandlerFunc {
	conf = c
	db = dtb
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}

// "/api" return configuration parameters
func ApiRoot(writer http.ResponseWriter, request *http.Request) {
	res, _ := json.Marshal(conf)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
	writer.Write([]byte("\n"))
}

func ApiGetWalls(writer http.ResponseWriter, request *http.Request) {

	walls, err := db.GetWalls()

	if err != nil || walls == nil {
		writer.WriteHeader(http.StatusNoContent)
		writer.Write([]byte(`{}`))
		return
	}

	res, _ := json.Marshal(walls)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
	writer.Write([]byte("\n"))
}
