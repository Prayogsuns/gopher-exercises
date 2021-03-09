package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/prayogsuns/gopher-exercises/urlshort"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

var yamlf *string
var jsonf *string
var dbopt *bool

func init() {
	yamlf = flag.String("yaml", "", "yaml input file")
	jsonf = flag.String("json", "", "json input file")
	dbopt = flag.Bool("db", false, "db input flag")
	flag.Parse()
}

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	/*
			yaml := `
		- path: /urlshort
		  url: https://github.com/gophercises/urlshort
		- path: /urlshort-final
		  url: https://github.com/gophercises/urlshort/tree/solution
		`
	*/
	if *yamlf != "" {
		yamlF, errry := ioutil.ReadFile(*yamlf)
		check(errry)
		//yaml := string(yamlF)

		//serverHandler, erry := urlshort.YAMLHandler([]byte(yaml), mapHandler)
		serverHandler, erry := urlshort.YAMLHandler(yamlF, mapHandler)
		check(erry)

		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", serverHandler)
	} else if *jsonf != "" {
		jsonF, errrj := ioutil.ReadFile(*jsonf)
		check(errrj)
		//json := string(jsonF)

		//serverHandler, errj := urlshort.JSONHandler([]byte(json), mapHandler)
		serverHandler, errj := urlshort.JSONHandler(jsonF, mapHandler)
		check(errj)

		fmt.Println("Starting the server on :8081")
		http.ListenAndServe(":8081", serverHandler)
	} else if *dbopt {
		fmt.Println("Db Actions")
    var DB_NAME = "sqllite3db"
    var filename = DB_NAME + ".db"
    var table = map[string]string{"name":"urlmap", "col1": "id", "col2": "path", "col3": "url"}
    
    db := createDb(filename) 
    createTable(db, table) 
    insertDataInDb(db, table) 

		queryDb(db, table)
		closeDb(filename, db)
		fmt.Println("Db Closed")
	}
	//http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createDb(filename string) *sql.DB {
  db, err := sql.Open("sqlite3", filename)
	check(err)
  return db
}

func createTable(db *sql.DB, table map[string]string) {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS " + table["name"] + " (" + table["col1"] + " INTEGER PRIMARY KEY, " + table["col2"] + " TEXT, " + table["col3"] + " TEXT)")
	check(err)
  stmt.Exec()
}

func insertDataInDb(db *sql.DB, table map[string]string) {
  stmt, err := db.Prepare("INSERT INTO " + table["name"] + " (" + table["col1"] +", " + table["col2"] + ", " + table["col3"] + ") VALUES (?, ?, ?)")
	check(err)
	stmt.Exec(001, "/urlshort", "https://github.com/gophercises/urlshort")
	stmt.Exec(002, "/urlshort-final", "https://github.com/gophercises/urlshort/tree/solution")
}

func queryDb(db *sql.DB, table map[string]string) {
  rows, err := db.Query("SELECT " + table["col1"] + ", " + table["col2"] + ", " + table["col3"] + " from " + table["name"]) 
	check(err)
	var id int
	var path string
	var url string
	for rows.Next() {
    rows.Scan(&id, &path, &url)
		fmt.Println(strconv.Itoa(id) + ": " + path + " " + url)
	}
}

func closeDb(filename string, db *sql.DB) {
	db.Close() 
  os.Remove(filename)
}
