package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var (
	HOST     = os.Getenv("DB_HOST")
	PORT     = os.Getenv("DB_PORT")
	USER     = os.Getenv("DB_USER")
	DBNAME   = os.Getenv("DB_NAME")
	PASSWORD = os.Getenv("DB_PASSWORD")
)

var (
	db    *sql.DB
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
