package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Signup handles signup requests.
func Signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rb := RequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&rb); err != nil {
		responseJSON(w, http.StatusBadRequest, ResponseError{
			Message: "Account creation failed",
			Cause:   "JSON is expected",
		})
		return
	}
	if err := rb.validateSignup(); err != nil {
		responseJSON(w, http.StatusBadRequest, ResponseError{
			Message: "Account creation failed",
			Cause:   err.Error(),
		})
		return
	}
	user := User{
		UserID:   rb.UserID,
		Password: rb.Password,
		Nickname: rb.UserID,
	}
	if err := user.create(); err != nil {
		responseJSON(w, http.StatusBadRequest, ResponseError{
			Message: "Account creation failed",
			Cause:   "already same user_id exists",
		})
		return
	}
	responseJSON(w, http.StatusOK, ResponseSignup{
		Message: "Account successfully created",
		User: User{
			UserID:   user.UserID,
			Nickname: user.Nickname,
		},
	})
	return
}
