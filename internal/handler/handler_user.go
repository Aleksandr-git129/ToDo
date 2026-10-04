package handler

import (
	"net/http"
	"todo-list/internal/model"
)

type HandlerUser struct {
	User *model.User
}

func HandlerUserConstructor(u *model.User) *HandlerUser {
	return &HandlerUser{
		User: u,
	}
}

func (h *HandlerUser) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *HandlerUser) GetUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *HandlerUser) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *HandlerUser) ChangeUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *HandlerUser) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

}
