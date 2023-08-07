package main

import "fmt"

type Feeds struct {
	TotalPosts int
	LatestPosts []Post
}

type Tag struct {
	Id int
	Name string
}

func(t Tag) TagUrl() string {
	return fmt.Sprintf("/blog/tag/%s", t.Name)
}

type Post struct {
	Id int
	Title string
	Slug string
	Body string
	Published string
	Tags []Tag
}

func(p *Post) AbsoluteUrl() string {
	return fmt.Sprintf("/blog/%d/%s", p.Id, p.Slug)
}

func(p *Post) CommentUrl() string {
	return fmt.Sprintf("/blog/%d/comment", p.Id)
}

type Comment struct {
	Id int
	Name string
	Body string
	Created string
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
	Feeds *Feeds
	TagName string
}

type DetailPost struct {
	Post Post
	Comments []Comment
	Feeds *Feeds
}
