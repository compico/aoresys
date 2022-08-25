package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/compico/aoresys/internal/controller"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err.Error())
	}

	connectToDB()
	defer db.Close()

	srv := &http.Server{
		Addr:         os.Getenv("ADDR"),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	router := httprouter.New()
	srv.Handler = router

	router.GET("/", controller.Middleware(controller.IndexController))
	router.GET("/doms/index", controller.Middleware(indexpageHandler))
	router.GET("/doms/registration", controller.Middleware(registrationHandler))
	router.GET("/doms/login", controller.Middleware(loginHandler))
	router.GET("/doms/servercard", controller.Middleware(servercardHandler))

	router.POST("/api/v1/register", controller.Middleware(registerApiHandler))
	router.POST("/api/v1/login", controller.Middleware(loginApiHandler))

	router.ServeFiles("/public/*filepath", http.Dir("./public/"))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err.Error())
	}
}
