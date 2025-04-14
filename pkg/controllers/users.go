package controllers

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ladecadence/PhotonAPI/pkg/models"
)

var ErrAuth = errors.New("Unauthorized")

func Authorize(request *http.Request) error {
	// check user
	username := request.Header.Get("X-User")
	if username == "" {
		return ErrAuth
	}
	user, err := db.GetUser(username)
	if err != nil || user.Name != username {
		return ErrAuth
	}

	// get cookie
	st, err := request.Cookie("session_token")
	if err != nil || st.Value == "" || st.Value != user.Token {
		return ErrAuth
	}

	// get CSRF token from headers
	csrf := request.Header.Get("X-CSRF-Token")
	if csrf == "" || csrf != user.CSRF {
		return ErrAuth
	}

	return nil
}

func ApiSignup(writer http.ResponseWriter, request *http.Request) {
	// check form fields
	username := request.FormValue("username")
	password := request.FormValue("password")
	email := request.FormValue("email")

	if len(username) > 3 && len(password) > 8 && len(email) > 6 {
		// user exists ?
		check, err := db.GetUser(username)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("{}\n"))
			return
		}
		if check.Name == username {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("{\"error\": \"existing user\"}\n"))
			return
		}
		// ok, create it
		user := models.User{
			Name:     username,
			Password: fmt.Sprintf("%x", sha256.Sum256([]byte(password))),
			Email:    email,
			Role:     models.UserRoleUser,
		}
		err = db.UpsertUser(user)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("{}\n"))
			return
		}

	} else {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("{\"error\": \"bad fields\"}\n"))
		return
	}
}

func ApiLogin(writer http.ResponseWriter, request *http.Request) {
	// check auth
	authOk := CheckAuth(request)
	if authOk {
		username, _, _ := request.BasicAuth()

		// generate tokens
		sessionToken := GenerateToken(32)
		csrfToken := GenerateToken(32)

		// generate cookies
		http.SetCookie(writer, &http.Cookie{
			Name:     "session_token",
			Value:    sessionToken,
			Expires:  time.Now().Add(7 * 24 * time.Hour),
			HttpOnly: true,
		})
		http.SetCookie(writer, &http.Cookie{
			Name:     "csrf_token",
			Value:    csrfToken,
			Expires:  time.Now().Add(7 * 24 * time.Hour),
			HttpOnly: false,
		})

		// store tokens
		user, err := db.GetUser(username)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("{}\n"))
			return
		}
		user.Token = sessionToken
		user.CSRF = csrfToken
		err = db.UpsertUser(user)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("{}\n"))
			return
		}
		writer.Write([]byte("{}\n"))

	} else {
		writer.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
	}
}

func ApiLogout(writer http.ResponseWriter, request *http.Request) {
	if err := Authorize(request); err != nil {
		writer.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
	}

	// clear cookies
	http.SetCookie(writer, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})
	http.SetCookie(writer, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: false,
	})

	// check user
	username := request.Header.Get("X-User")
	user, err := db.GetUser(username)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("{}\n"))
		return
	}
	user.CSRF = ""
	user.Token = ""
	db.UpsertUser(user)
	writer.Write([]byte("{}\n"))
}
