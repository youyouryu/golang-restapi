package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Update handles updarte requests.
func Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var (
		authID string
		err    error
	)
	if authID, err = authenticate(r); err != nil {
		responseJSON(w, http.StatusUnauthorized, ResponseError{
			Message: "Authentication Faild", // Faild -> failed (wrong test case)
		})
		return
	}
	userID := p.ByName("user_id")
	if userID != authID {
		responseJSON(w, http.StatusForbidden, ResponseError{
			Message: "No Permission for Update",
		})
		return
	}
	user, err := userByID(userID)
	if err != nil {
		responseJSON(w, http.StatusNotFound, ResponseError{
			Message: "No User found",
		})
		return
	}
	rb := RequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&rb); err != nil {
		responseJSON(w, http.StatusBadRequest, ResponseError{
			Message: "User updation failed",
			Cause:   "JSON is expected",
		})
		return
	}
	if err := rb.validateUpdate(); err != nil {
		responseJSON(w, http.StatusBadRequest, ResponseError{
			Message: "User updation failed",
			Cause:   err.Error(),
		})
		return
	}
	user.Nickname = rb.Nickname
	user.Comment = rb.Comment
	if err := user.update(); err != nil {
		responseJSON(w, http.StatusBadRequest, ResponseError{
			Message: "User updation failed",
			Cause:   err.Error(),
		})
		return
	}
	recipe := []User{
		{
			Nickname: user.Nickname,
			Comment:  user.Comment,
		},
	}
	responseJSON(w, http.StatusOK, ResponseUpdate{
		Message: "User successfully updated",
		Recipe:  recipe,
	})
	return
}
