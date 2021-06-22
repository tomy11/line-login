package db

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gopkg.in/mgo.v2"
)

type Connection interface {
	Close()
	DB() *mgo.Database
}

type conn struct {
	session *mgo.Session
}

func NewConnection() Connection {
	var c conn
	var err error
	url := getURL()
	c.session, err = mgo.Dial(url)
	if err != nil {
		log.Panicln(err.Error())
	}
	return &c
}

func (c *conn) Close() {
	c.session.Close()
}

func (c *conn) DB() *mgo.Database {
	return c.session.DB(os.Getenv("DB_NAME"))
}

func getURL() string {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Panicln("error load db port from environment", err.Error())
		port = 27017
	}
	// fmt.Printf("mongodb://%s:%s@%s:%d/%s",
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASS"),
	// 	os.Getenv("DB_HOST"),
	// 	port,
	// 	os.Getenv("DB_NAME"))

	return fmt.Sprintf("mongodb://%s:%d/%s",
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("DB_NAME"))
}
