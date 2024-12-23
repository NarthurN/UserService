package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type User struct {
	Login    string
	Password string
	Name     string
}

var users = map[string]User{}

func main() {
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/login", handleLogin)

	fmt.Println("Server is listening...")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		panic(err)
	}
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Не удалось прочитать запрос", http.StatusBadRequest)
		return
	}

	data := strings.Split(string(bodyBytes), " ")
	if len(data) != 3 {
		http.Error(w, "Нужен текст в виде: login password name", http.StatusBadRequest)
		return
	}
	login, password, name := data[0], data[1], data[2]
	if login == "" || password == "" || name == "" {
		http.Error(w, "login password name не должны быть пустыми", http.StatusBadRequest)
		return
	}

	token := fmt.Sprintf("%s", md5.Sum([]byte(login))) // получаем 265dsfsd54vd3434
	if _, ok := users[token]; ok {
		http.Error(w, "Такой пользователь уже существует", http.StatusForbidden)
		return
	}
	users[token] = User{
		Login:    login,
		Password: password,
		Name:     name,
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = fmt.Fprint(w, "Пользователь успешно создан")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Не удалось прочитать запрос", http.StatusBadRequest)
		return
	}
	data := strings.Split(string(bodyBytes), " ")
	if len(data) != 2 {
		http.Error(w, "Нужен текст в виде: login password", http.StatusBadRequest)
		return
	}
	login, password := data[0], data[1]
	for token, user := range users {
		if user.Login == login && user.Password == password {
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, "Пользователь успешно залогинился\n")
			_, _ = fmt.Fprintf(w, "Token %s", token)
			return
		}
		http.Error(w, "Нет такого пользователя", http.StatusBadRequest)
		return
	}
}
