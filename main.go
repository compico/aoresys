package main

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {

	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	testdata()
	router := httprouter.New()
	srv.Handler = router
	router.GET("/", indexHandler)
	router.GET("/login", loginHandler)
	router.GET("/register", registerHandler)
	router.ServeFiles("/public/*filepath", http.Dir("./public/"))

	log.Fatal(srv.ListenAndServe())
}
