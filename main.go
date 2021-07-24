package main

import (
	"context"
	"fmt"
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

	initDBClient()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := cdb.Client.Connect(ctx); err != nil {
		panic(fmt.Errorf("Error: %v\n", err.Error()))
	}
	defer func() {
		err := cdb.Client.Disconnect(ctx)
		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
		}
	}()

	router := httprouter.New()
	srv.Handler = router
	router.GET("/", indexHandler)

	router.GET("/doms/index", indexpageHandler)
	router.GET("/doms/loginreg", loginregHandler)
	router.POST("/doms/existusername", existUsernameHandler)

	router.POST("/api/v1/register", registerApiHandler)

	router.ServeFiles("/public/*filepath", http.Dir("./public/"))

	log.Fatal(srv.ListenAndServe())
}
