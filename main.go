package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

var db *sql.DB

func getParamString(param string, defaultValue string) string {
	env := os.Getenv(param)
	if env != "" {
		return env
	}
	return defaultValue
}

func getConnectionString() (cs string) {
	user := getParamString("MYSQL_USER", "root")
	pw := getParamString("MYSQL_ROOT_PASSWORD", "")
	host := getParamString("MYSQL_HOST", "localhost")
	port := getParamString("MYSQL_PORT", "3306")
	dbName := getParamString("MYSQL_DATABASE", "restapi")
	cs = fmt.Sprintf("%s:%s@tcp([%s]:%s)/%s", user, pw, host, port, dbName)
	return
}

func initDb() {
	var err error
	cs := getConnectionString()
	log.Println(cs)
	db, err = sql.Open("mysql", cs)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to mysql.")
}

func initRouter() (r *httprouter.Router) {
	r = httprouter.New()
	r.POST("/signup", Signup)
	r.GET("/users/:user_id", Show)
	r.PATCH("/users/:user_id", Update)
	r.POST("/close", Close)
	return
}

func main() {
	initDb()
	defer db.Close()

	router := initRouter()
	log.Println("Server started.")
	log.Fatal(http.ListenAndServe(":8080", router))
}
