package main

import (
	"encoding/json"
	"fmt"
	"github.com/domesama/golang-chi-router-helloworld/packages/newsfeed"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {

	router := chi.NewRouter()

	feed := newsfeed.New()
	feed.Add(newsfeed.Item{"OwO", "Post UwU"})

	router.Get("/", func(w http.ResponseWriter, r *http.Request){
		_,err := w.Write([]byte("Hello World"))
		if err != nil{
			log.Fatal(err)
		}
	})

	router.Get("/newsfeed", func(w http.ResponseWriter,req *http.Request){
		items := feed.GetAll()
		json.NewEncoder(w).Encode(items)
	})

	router.Post("/newsfeed", func(w http.ResponseWriter, req *http.Request){
		request := map[string]string{}
		json.NewDecoder(req.Body).Decode(&request)

		feed.Add(newsfeed.Item{Title: request["title"], Post: request["title"]})

		w.WriteHeader(200)
	})

	//http.ListenAndServe(":8080", router)
	server := &http.Server{
		Addr:             portNumber,
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	fmt.Println("Starting server on port ", portNumber)
	log.Fatal(server.ListenAndServe())
}
