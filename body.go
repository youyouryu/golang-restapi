package main

import (
	"errors"
	"regexp"
	"unicode/utf8"
)

// RequestBody represents of signup and update
type RequestBody struct {
	User
}

func (b RequestBody) validateSignup() error {
	if b.UserID == "" || b.Password == "" {
		return errors.New("required user_id and password")
	}
	if l := utf8.RuneCountInString(b.UserID); l < 6 || 20 < l {
		return errors.New("length of user_id must be 6 to 20 characters")
	}
	if l := utf8.RuneCountInString(b.Password); l < 8 || 20 < l {
		return errors.New("length of password must be 6 to 20 characters")
	}
	if r := `^[a-zA-Z0-9]{6,20}$`; !regexp.MustCompile(r).MatchString(b.UserID) {
		return errors.New("pattern of user_id must match '" + r + "'")
	}
	if r := `^[!-~]{8,20}$`; !regexp.MustCompile(r).MatchString(b.Password) {
		return errors.New("pattern of password must match '" + r + "'")
	}
	return nil
}

func (b RequestBody) validateUpdate() error {
	if b.Nickname == "" && b.Comment == "" {
		return errors.New("required nickname or comment")
	}
	if l := utf8.RuneCountInString(b.Nickname); 30 < l {
		return errors.New("length of nickname must be up to 30 characters")
	}
	if l := utf8.RuneCountInString(b.Nickname); 100 < l {
		return errors.New("length of comment must be up to 100 characters")
	}
	if r := `^[^\x00-\x1F\x7F]*$`; !regexp.MustCompile(r).MatchString(b.Nickname) {
		return errors.New("pattern of nickname must match '" + r + "'")
	}
	if r := `^[^\x00-\x1F\x7F]*$`; !regexp.MustCompile(r).MatchString(b.Comment) {
		return errors.New("pattern of comment must match '" + r + "'")
	}
	if b.UserID != "" || b.Password != "" {
		return errors.New("not updatable user_id and password")
	}
	return nil
}

// ResponseError represents a response body to any invalid request.
type ResponseError struct {
	Message string `json:"message"`
	Cause   string `json:"cause,omitempty"`
}

// ResponseSignup represents a response body to signup request.
type ResponseSignup struct {
	Message string `json:"message"`
	User    `json:"user"`
}

// ResponseShow represents a response body to show request.
type ResponseShow struct {
	Message string `json:"message"`
	User    `json:"user"`
}

// ResponseUpdate represents a response body to update request.
type ResponseUpdate struct {
	Message string `json:"message"`
	Recipe  []User `json:"recipe"`
}

// ResponseDelete represents a response body to delete request.
type ResponseDelete struct {
	Message string `json:"message"`
}
