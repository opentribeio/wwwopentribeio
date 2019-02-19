package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var header string = "" // Authz header

type Subject struct {
	Data []struct {
		ID        int         `json:"ID"`
		CreatedAt time.Time   `json:"CreatedAt"`
		UpdatedAt time.Time   `json:"UpdatedAt"`
		DeletedAt interface{} `json:"DeletedAt"`
		Name      string      `json:"name"`
		Value     int         `json:"value"`
		Avg       string      `json:"avg"`
		UserID    int         `json:"user_id"`
	} `json:"data"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

// login struct
type Response struct {
	Account struct {
		ID        int         `json:"ID"`
		CreatedAt time.Time   `json:"CreatedAt"`
		UpdatedAt time.Time   `json:"UpdatedAt"`
		DeletedAt interface{} `json:"DeletedAt"`
		Email     string      `json:"email"`
		Password  string      `json:"password"`
		Token     string      `json:"token"`
	} `json:"account"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Print(r.Form)
	u := strings.Join(r.Form["email"], ", ")
	p := strings.Join(r.Form["pass"], ", ")
	url := "http://localhost:8000/api/user/new"
	data := Response{}
	body := post(u, p, url)

	json.Unmarshal(body, &data)
	if data.Status == true {
		http.Redirect(w, r, "/profile", http.StatusTemporaryRedirect)
	} else {
		http.Redirect(w, r, "/login/signup.html", http.StatusTemporaryRedirect)
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("html/login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		u := strings.Join(r.Form["username"], ", ")
		p := strings.Join(r.Form["password"], ", ")
		url := "http://localhost:8000/api/user/login"
		body := post(u, p, url)

		data := Response{}
		json.Unmarshal(body, &data)

		// If successful login
		if data.Status == true {
			header = "Bearer " + data.Account.Token //set auth token
			http.Redirect(w, r, "/profile", http.StatusTemporaryRedirect)
		} else {
			http.Redirect(w, r, "/login/login.html", http.StatusTemporaryRedirect)
		}
	}
}

func post(username string, password string, url string) []byte {
	//url := "http://localhost:8000/api/user/login"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{"email":"` + username + `", "password" : "` + password + `"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	return body

}
func get(url string) []byte {
	fmt.Println("URL:>", url)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", header)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
func Profile(w http.ResponseWriter, r *http.Request) {
	body_me := get("http://localhost:8000/api/me/tax")
	body_global := get("http://localhost:8000/api/tribe/stats")

	data := Subject{}
	json.Unmarshal(body_me, &data)

	d := Subject{}
	json.Unmarshal(body_global, &d)
	//	json.Unmarshal(body_global, &data)
	for i, _ := range data.Data {
		data.Data[i].Avg = d.Data[i].Avg
	}

	tmpl, err := template.ParseFiles("html/profile.html")
	if err != nil {
		fmt.Print(err)
	}
	tmpl.Execute(w, data)
}
