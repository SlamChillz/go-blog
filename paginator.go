package main

import (
	"net/http"
	"strings"
	"strconv"

	"github.com/gorilla/mux"
)

const PageLimit = 3

func Paginate(r *http.Request) (*Pager, int, error) {
	var page int
	var offset int
	var err error
	var pager *Pager
	var count int
	vars := mux.Vars(r)
	queryValues := r.URL.Query()
	page, err = strconv.Atoi(strings.TrimSpace(queryValues.Get("page")))
	if err != nil {
		page = 1
	}
	if page <= 0 {
		page = 1
	}
	if tagName, ok := vars["name"]; ok {
		count, err = CountPostsByTag(strings.TrimSpace(tagName))
	} else {
		count, err = CountPosts()
	}
	if err != nil {
		return pager, offset, err
	}
	totalPages := count / PageLimit
	if (count % PageLimit) > 0 {
		totalPages += 1
	}
	if page > totalPages {
		page = totalPages
	}
	offset = (page - 1) * PageLimit
	pager = &Pager{pageNumber: page, totalPages: totalPages}
	return pager, offset, nil
}
