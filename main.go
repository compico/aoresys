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

	testdata()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := cdb.Client.Connect(ctx)
	if err != nil {
		panic(fmt.Errorf("Error: %v\n", err.Error()))
	}
	defer func() {
		err := cdb.Client.Disconnect(ctx)
		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
		}
	}()

	// conf, err := mail.LoadConfig("./conf/smtp.json")
	// if err != nil {
	// 	fmt.Printf("Error:%v\n", err.Error())
	// }
	// mailclient := mail.NewClient()
	// mailclient.NewAuth(conf)
	// msg := mail.NewMessage()
	// msg.AddTo(
	// 	"compico@mail.ru",
	// 	"ysofe.may0@arktive.com",
	// 	"x10mahdi10d@saymuscge.ml",
	// 	"qpaint@marchmovo.com",
	// )
	// msg.AddSubject("Hello world!")
	// msg.AddMessage("Test mail shtuki\nAndTestSlashN")
	// err = msg.CompileMail()
	// if err != nil {
	// 	fmt.Printf("Error:%v\n", err.Error())
	// }
	// err = mailclient.SendMail(*msg)
	// if err != nil {
	// 	fmt.Printf("Error:%v\n", err.Error())
	// }

	router := httprouter.New()
	srv.Handler = router

	router.GET("/", indexHandler)
	router.GET("/login", loginHandler)
	router.GET("/register", registerHandler)

	router.POST("/api/v1/register", registerApiHandler)
	router.GET("/api/v1/existusername/:username", existUsernameHandler)
	router.ServeFiles("/public/*filepath", http.Dir("./public/"))

	log.Fatal(srv.ListenAndServe())
}
