package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type User struct {
	Username   string
	Email      string
	Password   string
	Name       string
	Surname    string
	Patronymic string
}

func GenPasswd() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 10
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}

func GenData() User {
	fullName := strings.Split(randomdata.FullName(randomdata.RandomGender), " ")
	user := User{
		Username:   randomdata.SillyName(),
		Email:      randomdata.Email(),
		Password:   GenPasswd(),
		Name:       fullName[0],
		Surname:    fullName[1],
		Patronymic: "",
	}
	return user
}

func CreateUser(user User) {
	url := "http://127.0.0.1:8080/user"
	body, _ := json.Marshal(user)

	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ = ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

}

func main() {
	for i := 0; i < 10000; i++ {
		data := GenData()
		fmt.Println(data)
		CreateUser(data)
	}
}
