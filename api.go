package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func registerApiHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// name
	// password
	// passwordConfirmation
	// email
}

func existUsernameHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	username := ps.ByName("username")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	res, err := cdb.ExistUsername(ctx, username)
	if err != nil {
		fmt.Fprintf(w, "{ \"error\":\"%v\" }", err)
		return
	}
	fmt.Fprintf(w, "{ \"existuser\" : %t }", res)
}
