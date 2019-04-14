package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var templates *template.Template

func main() {

	templates = template.Must(template.ParseGlob("static/*.html"))

	fmt.Println("Running...")
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", getHandler).Methods("GET")
	r.HandleFunc("/", postHandler).Methods("POST")
	r.HandleFunc("/", putHandler).Methods("PUT")
	r.HandleFunc("/", deleteHandler).Methods("DELETE")
	r.HandleFunc("/admin", adminHandler).Methods("POST")

	r.HandleFunc("/index", index).Methods("GET")

	r.HandleFunc("/json", jsonGetHandler).Methods("GET")
	r.HandleFunc("/json", jsonPostHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(b))
}

func info(r *http.Request) string {
	return fmt.Sprintf("%s/ %s %s%s", r.Method, r.Proto, r.Host, r.URL)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	pretty(keys)
	fmt.Fprint(w, info(r))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	pretty(keys)
	r.ParseForm()
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	fmt.Println(email, password)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, info(r))
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	pretty(keys)
	r.ParseForm()
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	fmt.Println(email, password)
	fmt.Fprint(w, info(r))
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	pretty(keys)
	fmt.Fprint(w, info(r))
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	if email == "admin" && password == "admin" {
		fmt.Fprint(w, info(r))
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprint(w, http.StatusText(http.StatusUnauthorized))
}

func index(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func jsonGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF8")
	json.NewEncoder(w).Encode(struct {
		Method      string `json:"method"`
		HttpVersion string `json:"http_version"`
		Host        string `json:"host"`
		Url         string `json:"url"`
	}{
		Method:      fmt.Sprintf("%s", r.Method),
		HttpVersion: fmt.Sprintf("%s", r.Proto),
		Host:        fmt.Sprintf("%s", r.Host),
		Url:         fmt.Sprintf("%s", r.URL),
	})
}

type JsonData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func jsonPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	data := JsonData{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(struct {
			Info       string `json:"info"`
			StatusCode int    `json:"status_code"`
			StatusText string `json:"status_text"`
			Error      string `json:"error"`
		}{
			Info:       info(r),
			StatusCode: http.StatusUnprocessableEntity,
			StatusText: http.StatusText(http.StatusUnprocessableEntity),
			Error:      err.Error(),
		})
		return
	}

	fmt.Println(data)

	if data.Email == "admin" && data.Password == "admin" {
		json.NewEncoder(w).Encode(struct {
			Method      string   `json:"method"`
			HttpVersion string   `json:"http_version"`
			Host        string   `json:"host"`
			Url         string   `json:"url"`
			Data        JsonData `json:"data"`
		}{
			Method:      fmt.Sprintf("%s", r.Method),
			HttpVersion: fmt.Sprintf("%s", r.Proto),
			Host:        fmt.Sprintf("%s", r.Host),
			Url:         fmt.Sprintf("%s", r.URL),
			Data:        data,
		})
		return
	}

	w.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(w).Encode(struct {
		Info       string `json:"info"`
		StatusCode int    `json:"status_code"`
		StatusText string `json:"status_text"`
	}{
		Info:       info(r),
		StatusCode: http.StatusUnprocessableEntity,
		StatusText: http.StatusText(http.StatusUnprocessableEntity),
	})
}
