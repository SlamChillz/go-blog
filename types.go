package main

import "fmt"

type Post struct {
	Id int
	Title string
	Slug string
	Body string
	Published string
}

func(p *Post) AbsoluteUrl() string {
	return fmt.Sprintf("/blog/%d/%s", p.Id, p.Slug)
}


type Pager struct {
	pageNumber int
	totalPages int
}

func(p *Pager) Page() int {
	return p.pageNumber
}

func(p *Pager) HasPrev() bool {
	if p.pageNumber > 1 {
		return true
	}
	return false
}

func(p *Pager) Prev() int {
	return p.pageNumber - 1
}

func(p *Pager) HasNext() bool {
	if p.pageNumber < p.totalPages {
		return true
	}
	return false
}

func(p *Pager) Next() int {
	return p.pageNumber + 1
}

func(p *Pager) TotalPages() int {
	return p.totalPages
}


type ListPosts struct {
	Posts []Post
	Pager *Pager
}
