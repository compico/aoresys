package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

var tpath = "./templates/"

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles(
		tpath+"head.html",
		tpath+"index.html",
		tpath+"footer.html",
	)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "index", nil)
	if err != nil {
		fmt.Fprintf(w, "[ERROR] %v!!", err.Error())
		fmt.Printf("[ERROR] %v!!", err.Error())
	}
}

func loginregHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("HX-Trigger", "{\"addBtnEventStatements\": \"\"}")
	t, err := template.ParseFiles(tpath + "loginreg.html")
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "loginreg", nil)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err.Error())
	}
}

func indexpageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles(tpath + "indexpage.html")
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "indexpage", nil)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err.Error())
	}
}

func existUsernameHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	res := struct {
		Result   bool
		Username string
	}{
		Username: r.FormValue("username"),
	}
	var err error
	if res.Username != "" {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		res.Result, err = cdb.ExistUsername(ctx, res.Username)
		if err != nil {
			fmt.Fprintf(w, "{ \"error\":\"%v\" }", err)
			return
		}
	}
	t, err := template.ParseFiles(tpath + "existusername.html")
	if err != nil {
		fmt.Fprintf(w, "{ \"error\":\"%v\" }", err)
		return
	}
	err = t.ExecuteTemplate(w, "existusername", res)
	if err != nil {
		fmt.Fprintf(w, "{ \"error\":\"%v\" }", err)
		return
	}
}
