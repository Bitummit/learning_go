package user

import (
	"net/http";
	"github.com/julienschmidt/httprouter";
	"go_api/internal/handlers"
)

const (
	usersURL = "/users"
	userURL = "/users/:uuid"
)

type handler struct {

}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler)Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetUserList)
	router.GET(userURL, h.GetUserById)
	router.POST(usersURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartUpdateUser)
	router.DELETE(usersURL, h.DeleteUser)

}

func (h *handler) GetUserList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("user list!"))
}

func (h *handler) GetUserById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("user info"))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(201)
	w.Write([]byte("user create!"))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("user update!"))
}

func (h *handler) PartUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("user part update!"))
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("user delete!"))
}

