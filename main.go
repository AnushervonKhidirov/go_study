package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type FetchResult[Type any] struct {
	data Type
	err  error
}

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Address  Address `json:"address"`
	Company  Company `json:"company"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

const userEndpoint = "https://jsonplaceholder.typicode.com/users"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	users, err := fetchUsers(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(users)
}

func fetchUsers(ctx context.Context) ([]User, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, userEndpoint, http.NoBody)

	if err != nil {
		return nil, err
	}

	resCh := make(chan FetchResult[[]User], 1)

	go func() {
		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			resCh <- FetchResult[[]User]{err: err}
			return
		}

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			resCh <- FetchResult[[]User]{err: err}
			return
		}

		var users []User

		err = json.Unmarshal(body, &users)

		if err != nil {
			resCh <- FetchResult[[]User]{err: err}
			return
		}

		resCh <- FetchResult[[]User]{data: users}
	}()

	select {
	case res := <-resCh:
		if res.err != nil {
			return nil, res.err
		}

		return res.data, nil

	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
