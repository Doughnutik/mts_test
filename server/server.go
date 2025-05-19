package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	gen "mts_test/gen"
)

type userService struct {
	users map[string]string
	mux   sync.Mutex
}

func (service *userService) RegisterPost(ctx context.Context, req *gen.AuthData) (gen.RegisterPostRes, error) {
	service.mux.Lock()
	defer service.mux.Unlock()
	_, ok := service.users[req.GetEmail()]
	if !ok {
		service.users[req.GetEmail()] = req.GetPassword()
		result := &gen.AuthTokenResponse{}
		result.Token.SetTo(fmt.Sprintf("token+%s+%s", req.Email, req.Password))
		return result, nil
	} else {
		return &gen.RegisterPostConflict{Error: gen.NewOptString("email уже существует")}, nil
	}
}

func main() {
	service := &userService{
		users: map[string]string{},
	}
	srv, err := gen.NewServer(service)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
