package main

import (
	"go_api/internal/user"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)


func main() {
	
	log.Println("Create router")
	router := httprouter.New()

	handler := user.NewHandler()
	handler.Register(router)

	log.Println("Register user handler")
	start(router)
	
}

func start(router *httprouter.Router) {
	log.Println("Start application ...")
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler: router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	log.Println("server is listening 0.0.0.0:8000")
	log.Fatalln(server.Serve(listener))
}