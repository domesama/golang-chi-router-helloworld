package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {

	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request){
		_,err := w.Write([]byte("Hello World"))
		if err != nil{
			log.Fatal(err)
		}
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
