package main

import (
	"database/sql"
	"fmt"
	"os"
	"log"
	_ "github.com/lib/pq"
)

var (
	HOST = os.Getenv("BLOG_HOST")
	PORT = os.Getenv("BLOG_PORT")
	USER = os.Getenv("BLOG_USER")
	DBNAME = os.Getenv("BLOG_DBNAME")
	PASSWORD = os.Getenv("BLOG_DBPASSWORD")
)

var (
	db *sql.DB
	feeds *Feeds
)


func init() {
	var err error
	connInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DBNAME)
	db, err = sql.Open("postgres", connInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	bootStrapFeeds()
	fmt.Printf("Connected to database successfully.\n")
}

func bootStrapFeeds() {
	var err error
	feeds = &Feeds{}
	feeds.TotalPosts, err = CountPosts()
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	feeds.LatestPosts, err = GetAllPosts(0)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
}

// BLOG_HOST=localhost BLOG_PORT=5432 BLOG_USER=mendy BLOG_DBNAME=blog BLOG_DBPASSWORD=mendy