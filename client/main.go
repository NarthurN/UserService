package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const baseURL = "http://localhost:9000"

func main() {
	if err := register("user1", "user1", "user1"); err != nil {
		panic(err)
	}
	if err := login("user1", "user1"); err != nil {
		panic(err)
	}
}

func register(login, password, name string) error {
	fullUrl := baseURL + "/register"
	requestBody := fmt.Sprintf("%s %s %s", login, password, name)

	request, err := http.NewRequest(http.MethodPost, fullUrl, strings.NewReader(requestBody))
	if err != nil {
		return err
	}
	
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(responseBody))

	return nil
}

func login(login, password string) error {
	fullUrl := baseURL + "/login"
	requestBody := fmt.Sprintf("%s %s", login, password)

	request, err := http.NewRequest(http.MethodPost, fullUrl, strings.NewReader(requestBody))
	if err != nil {
		return err
	}
	
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(responseBody))

	return nil	
}