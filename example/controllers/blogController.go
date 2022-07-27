package controllers

import "net/http"

func AllBlog(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("AllBlog"))
}

func GetBlog(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("GetBlog"))
}

func PostBlog(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("PostBlog"))
}

func EditBlog(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("EditBlog"))
}

func DelBlog(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("DelBlog"))
}
