package controllers

import (
	"net/http"
	"regexp"
)
type userController struct{
	userIdPattern *regexp.Regexp
}

func (uc userController ) ServerHTTP (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome to the User Controller"))
}

func newUserControler () *userController{
	return &userController{
		userIdPattern: regexp.MustCompile(`^/users/(\d+)/?`)
	}
}