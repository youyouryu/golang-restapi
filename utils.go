package main

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
)

func responseJSON(w http.ResponseWriter, status int, body interface{}) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(jsonBytes)
	return
}

func passwordHash(password string) (hash string) {
	b := sha512.Sum512([]byte(password))
	hash = hex.EncodeToString(b[:])
	return
}

func authenticate(r *http.Request) (userID string, err error) {
	userID, password, ok := r.BasicAuth()
	if !ok {
		err = errors.New("could not parse Authorization header")
		return
	}
	user, err := userByID(userID)
	if err != nil {
		err = errors.New("no user found")
		return
	}
	if passwordHash(password) != user.Password {
		err = errors.New("password does not match")
		return
	}
	return
}
