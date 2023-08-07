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
	var detailPost = DetailPost{}
	vars := mux.Vars(r)
	slug := vars["slug"]
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Bad request: Invalid blog id.", http.StatusBadRequest)
		return
	}
	t, err := template.ParseFiles("templates/detail.html")
	if err == nil {
		detailPost.Post, err = GetSinglePost(id, slug)
		if err == nil {
			detailPost.Comments, err = AllPostComments(id)
			if err == nil {
				err = t.Execute(w, &detailPost)
			}
		} else if err == sql.ErrNoRows {
			Error(w, http.StatusNotFound, err)
			return
		}
	}
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
	}
}

func PostComment(w http.ResponseWriter, r *http.Request) {
	var t *template.Template
	vars := mux.Vars(r)
	comment := map[string]interface{} {
		"success": false,
		"post": &Post{},
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Bad request: Invalid blog id.", http.StatusBadRequest)
		return
	}
	t, err = template.ParseFiles("templates/comment.html")
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}
	comment["post"], err = GetPostById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			Error(w, http.StatusNotFound, err)
		} else {
			Error(w, http.StatusInternalServerError, err)
		}
		return
	}
	name := r.PostFormValue("name")
	body := r.PostFormValue("body")
	if (len(name) > 0) && (len(body) > 0) {
		err = CreateComment(id, name, body)
		if err == nil {
			comment["success"] = true
		}
	}
	err = t.Execute(w, &comment)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
	}
}
