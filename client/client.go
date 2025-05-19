package main

import (
	"context"
	"fmt"

	gen "mts_test/gen"
)

func run(ctx context.Context) error {
	url := "http://localhost:8080"
	var emails []string = []string{"email@first", "email@second", "email@first"}
	var passwords []string = []string{"first", "second", "first"}

	client, err := gen.NewClient(url)
	if err != nil {
		return fmt.Errorf("create client: %w", err)
	}

	for i, email := range emails {
		res, err := client.RegisterPost(ctx, &gen.AuthData{
			Email:    email,
			Password: passwords[i],
		})
		if err != nil {
			return fmt.Errorf("ошибка запроса регистрации с email = %s и password = %s: %w",
				email, passwords[i], err)
		}

		switch p := res.(type) {
		case *gen.AuthTokenResponse:
			fmt.Printf("token = %s\n", p.Token.Value)
		case *gen.RegisterPostConflict:
			fmt.Println("email уже существует")
		}
	}
	return nil
}

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Println(err)
	}
}

// curl -X 'POST' \
//   'http://localhost:8080/register' \
//   -H 'accept: application/json' \
//   -H 'Content-Type: application/json' \
//   -d '{
//   "email": "email@email",
//   "password": "password"
// }'
