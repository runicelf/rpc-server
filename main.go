package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/runicelf/rpc-server/config"
	"github.com/runicelf/rpc-server/controller"
	"github.com/runicelf/rpc-server/repository"
	"net/http"
)

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")

	db := repository.New(config.Get())
	err := s.RegisterService(controller.New(db), "")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.Handle("/rpc", s)

	fmt.Println("service started...")
	err = http.ListenAndServe(":9999", r)
	if err != nil {
		panic(err)
	}
}
