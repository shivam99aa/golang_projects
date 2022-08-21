package controllers

import (
	"fmt"
	"golang_projects/photogallery_app/views"
	"net/http"
)

type Users struct {
	NewView *views.View
}

type SignUpForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/signup"),
	}
}

// Create is used to process the signup form when a user submits it.
// This is used to create a new user account.
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	signUpForm := new(SignUpForm)

	if err := ParseForm(r, signUpForm); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, signUpForm)
}
