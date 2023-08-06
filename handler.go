package main

import (
	"net/http"
	"html/template"
	"log"
	"database/sql"
	"strconv"

	"github.com/gorilla/mux"
)

func Error(w http.ResponseWriter, code int, err error) {
	go log.Printf("%v", err)
	http.Error(w, http.StatusText(code), code)
}

func PostList(w http.ResponseWriter, r *http.Request) {
	var err error
	var offset int
	var t *template.Template
	listPosts := ListPosts{}
	listPosts.Pager, offset, err = Paginate(r)
	if err == nil {
		t, err = template.ParseFiles("templates/list.html")
		if err == nil {
			listPosts.Posts, err = GetAllPosts(offset)
			if err == nil {
				err = t.Execute(w, &listPosts)
			}
		}
	}
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
	}
}

func PostDetail(w http.ResponseWriter, r *http.Request) {
	var err error
	var post Post
	vars := mux.Vars(r)
	slug := vars["slug"]
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Bad request: Invalid blog id.", http.StatusBadRequest)
		return
	}
	t, err := template.ParseFiles("templates/detail.html")
	if err == nil {
		post, err = GetSinglePost(id, slug)
		if err == nil {
			err = t.Execute(w, &post)
		} else if err == sql.ErrNoRows {
			Error(w, http.StatusNotFound, err)
			return
		}
	}
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
	}
}
