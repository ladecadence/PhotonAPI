package controllers

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/ladecadence/PhotonAPI/pkg/config"
	"github.com/ladecadence/PhotonAPI/pkg/database"
	"github.com/ladecadence/PhotonAPI/pkg/models"
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

func CheckAuth(r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if ok {
		// get username from DB
		user, err := db.GetUser(username)
		if err != nil {
			return false
		}
		// ok, we have a username, check password
		passwordHash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
		passwordMatch := (subtle.ConstantTimeCompare([]byte(passwordHash), []byte(user.Password)) == 1)
		if passwordMatch {
			return true
		} else {
			return false
		}

	} else {
		return false
	}
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
		data, err := json.Marshal(wall)
		writer.WriteHeader(http.StatusOK)
		writer.Write(data)
		writer.Write([]byte("\n"))
	} else {
		writer.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
	}
}

func ApiGetProblems(writer http.ResponseWriter, request *http.Request) {

	problems, err := db.GetProblems()

	if err != nil || problems == nil {
		writer.WriteHeader(http.StatusNoContent)
		writer.Write([]byte(`{}\n`))
		return
	}

	res, _ := json.Marshal(problems)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
	writer.Write([]byte("\n"))
}

func ApiGetProblem(writer http.ResponseWriter, request *http.Request) {
	// get ID
	uid := request.PathValue("uid")
	if uid == "" {
		writer.WriteHeader(http.StatusNoContent)
		writer.Write([]byte(`{}\n`))
		return
	}
	problem, err := db.GetProblem(uid)

	if err != nil {
		writer.WriteHeader(http.StatusNoContent)
		writer.Write([]byte(`{}\n`))
		return
	}

	res, _ := json.Marshal(problem)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
	writer.Write([]byte("\n"))
}

func ApiNewProblem(writer http.ResponseWriter, request *http.Request) {
	// check auth
	authOk := CheckAuth(request)
	if authOk {
		reqBody, _ := io.ReadAll(request.Body)
		request.Body.Close()
		// try to create new wall
		problem := models.Problem{}
		err := json.Unmarshal(reqBody, &problem)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			log.Printf("❌ Error decoding body: %v", err.Error())
			writer.Write([]byte("{}\n"))
			return
		}
		err = db.UpsertProblem(problem)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("{}\n"))
			return
		}
		data, err := json.Marshal(problem)
		writer.WriteHeader(http.StatusOK)
		writer.Write(data)
	} else {
		writer.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
	}
}

func ApiGetWallProblems(writer http.ResponseWriter, request *http.Request) {
	// get ID
	wallid := request.PathValue("wallid")
	if wallid == "" {
		writer.WriteHeader(http.StatusNoContent)
		writer.Write([]byte(`{}\n`))
		return
	}
	problem, err := db.GetWallProblems(wallid)

	if err != nil {
		writer.WriteHeader(http.StatusNoContent)
		writer.Write([]byte(`{}\n`))
		return
	}

	res, _ := json.Marshal(problem)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
	writer.Write([]byte("\n"))
}
