package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/ladecadence/PhotonAPI/pkg/models"
)

func ApiGetProblems(writer http.ResponseWriter, request *http.Request) {
	// check query
	page, _ := strconv.Atoi(request.URL.Query().Get("page"))
	page_size, _ := strconv.Atoi(request.URL.Query().Get("page_size"))

	// filter
	filter := models.ProblemFilter{Active: false}

	wallid := request.URL.Query().Get("wallid")
	grade_max := request.URL.Query().Get("grade_max")
	grade_min := request.URL.Query().Get("grade_min")
	order_by := request.URL.Query().Get("order_by")
	order_dir := request.URL.Query().Get("order_dir")

	if wallid != "" {
		filter.SetWallID(wallid)
		fmt.Println("Wallid: ", wallid)
	}
	if grade_max != "" || grade_min != "" {
		min, _ := strconv.Atoi(grade_min)
		max, _ := strconv.Atoi(grade_max)
		filter.SetGradeRange(min, max)
		fmt.Printf("Grade range: %v - %v\n", min, max)
	}
	if order_by != "" {
		switch order_by {
		case "grade":
			filter.SetOrderBy(models.FilterOrderByGrade)
		case "sends":
			filter.SetOrderBy(models.FilterOrderByGrade)
		case "name":
			filter.SetOrderBy(models.FilterOrderByName)
		default:
			filter.SetOrderBy(models.FilterOrderByNothing)
		}
	}
	if order_dir != "" {
		switch order_dir {
		case "asc":
			filter.SetOrderDir(models.FilterOrderAsc)
		case "desc":
			filter.SetOrderDir(models.FilterOrderDesc)
		default:
			filter.SetOrderDir(models.FilterOrderAsc)
		}
	}

	problems, err := db.GetProblems(page, page_size, filter)

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
		data, _ := json.Marshal(problem)
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
