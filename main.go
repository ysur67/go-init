package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	BS          string `json:"bs"`
}

type User struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
}

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func getJsonResponse(url string, output interface{}) {
	request, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(request.Body)
	err = decoder.Decode(&output)
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}
	usersData := []User{}
	getJsonResponse(os.Getenv("user-source"), &usersData)
	postsData := []Post{}
	getJsonResponse(os.Getenv("post-source"), &postsData)
	fmt.Println(usersData)
	fmt.Println(postsData)
}
