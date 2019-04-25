package controllers

import "net/http"

func GetHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Hello!"))
}

func PostHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Post Hello!"))
}

func PutHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Put Hello!"))
}

func PatchHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Patch Hello!"))
}

func DeleteHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Hello!"))
}
