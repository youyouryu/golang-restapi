package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Close handles close request.
func Close(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		userID string
		err    error
	)
	if userID, err = authenticate(r); err != nil {
		responseJSON(w, http.StatusUnauthorized, ResponseError{
			Message: "Authentication Faild",
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
	if err := user.delete(); err != nil {
		responseJSON(w, http.StatusNotFound, ResponseError{
			Message: "No User found",
		})
		return
	}
	responseJSON(w, http.StatusOK, ResponseDelete{
		Message: "Account and user successfully removed",
	})
	return
}
