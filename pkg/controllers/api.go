package controllers

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

func GenerateToken(lenght int) string {
	bytes := make([]byte, lenght)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatal("Failed to generate Token %v", err)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}

// "/api" return configuration parameters
func ApiRoot(writer http.ResponseWriter, request *http.Request) {
	res, _ := json.Marshal(conf)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
	writer.Write([]byte("\n"))
}

func ApiGetProblems(writer http.ResponseWriter, request *http.Request) {
	// check query
	page, _ := strconv.Atoi(request.URL.Query().Get("page"))
	page_size, _ := strconv.Atoi(request.URL.Query().Get("page_size"))

	problems, err := db.GetProblems(page, page_size, models.ProblemFilter{Active: false})

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
