package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	//"os"
	"github.com/prayogsuns/gopher-exercises/urlshort"
)

var yamlf *string
var jsonf *string

func init() {
	yamlf = flag.String("yaml", "", "yaml input file")
	jsonf = flag.String("json", "", "json input file")
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
	} else {
		jsonF, errrj := ioutil.ReadFile(*jsonf)
		check(errrj)
		//json := string(jsonF)

		//serverHandler, errj := urlshort.JSONHandler([]byte(json), mapHandler)
		serverHandler, errj := urlshort.JSONHandler(jsonF, mapHandler)
		check(errj)

		fmt.Println("Starting the server on :8081")
		http.ListenAndServe(":8081", serverHandler)
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
