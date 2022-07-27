package controllers

import "net/http"

func AllPost(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("AllPost"))
}

func GetPost(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("GetPost"))
}

func PostPost(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("PostPost"))
}

func EditPost(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("EditPost"))
}

func DelPost(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("DelPost"))
}
