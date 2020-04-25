package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Show handles show requests.
func Show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if _, err := authenticate(r); err != nil {
		responseJSON(w, http.StatusUnauthorized, ResponseError{
			Message: "Authentication Faild",
		})
		return
	}
	userID := p.ByName("user_id")
	user, err := userByID(userID)
	if err != nil {
		responseJSON(w, http.StatusNotFound, ResponseError{
			Message: "No User found",
		})
		return
	}
	responseJSON(w, http.StatusOK, ResponseShow{
		Message: "User details by user_id",
		User: User{
			UserID:   user.UserID,
			Nickname: user.Nickname,
			Comment:  user.Comment,
		},
	})
	return
}
