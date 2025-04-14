package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/ladecadence/PhotonAPI/pkg/models"
)

func ApiGetWalls(writer http.ResponseWriter, request *http.Request) {

	// check query
	queryFields := strings.Split(request.URL.Query().Get("fields"), ",")
	fields := []string{}

	if len(queryFields) > 0 {
		// check wich fields are real DB fields
		for _, f := range queryFields {
			if slices.Contains(models.WallFields(), f) {
				fields = append(fields, f)
			}
		}
	}

	walls, err := db.GetWalls(fields)

	if err != nil || walls == nil {
		writer.WriteHeader(http.StatusNoContent)
		writer.Write([]byte(`{}\n`))
		return
	}

	res, _ := json.Marshal(walls)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
	writer.Write([]byte("\n"))
}

func ApiGetWall(writer http.ResponseWriter, request *http.Request) {

	// get ID
	uid := request.PathValue("uid")
	if uid == "" {
		writer.WriteHeader(http.StatusNoContent)
		writer.Write([]byte(`{}\n`))
		return
	}
	wall, err := db.GetWall(uid)

	if err != nil {
		writer.WriteHeader(http.StatusNoContent)
		writer.Write([]byte(`{}\n`))
		return
	}

	res, _ := json.Marshal(wall)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
	writer.Write([]byte("\n"))
}

func ApiNewWall(writer http.ResponseWriter, request *http.Request) {
	// check auth
	authOk := CheckAuth(request)
	if authOk {
		reqBody, _ := io.ReadAll(request.Body)
		request.Body.Close()
		// try to create new wall
		wall := models.Wall{}
		err := json.Unmarshal(reqBody, &wall)
		if err != nil {
			log.Printf("❌ Error decoding body: %v", err.Error())
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("{}\n"))
			return
		}
		err = db.UpsertWall(wall)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("{}\n"))
			return
		}
		data, _ := json.Marshal(wall)
		writer.WriteHeader(http.StatusOK)
		writer.Write(data)
		writer.Write([]byte("\n"))
	} else {
		writer.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
	}
}
