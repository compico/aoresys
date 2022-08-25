package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpath = "./templates/"

type Navbar struct {
	Homebutton         string
	Registrationbutton string
	Loginbutton        string
}

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

func registrationHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// w.Header().Add("HX-Trigger", "{\"addBtnEventStatements\": \"example\"}")
	data := struct {
		Navbar Navbar
	}{
		Navbar{
			Registrationbutton: "active",
		},
	}
	t, err := template.ParseFiles(tpath+"registration.html", tpath+"navbar.html")
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "registration", data)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err.Error())
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := struct {
		Navbar Navbar
	}{
		Navbar{
			Loginbutton: "active",
		},
	}
	t, err := template.ParseFiles(tpath+"login.html", tpath+"navbar.html")
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "login", data)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err.Error())
	}
}

func servercardHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	dataservercard := getServerCardData("")
	t, err := template.ParseFiles(tpath + "servercard.html")
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "servercard", dataservercard)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err.Error())
	}
}

func indexpageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := struct {
		Navbar Navbar
	}{
		Navbar{
			Homebutton: "active",
		},
	}
	t, err := template.ParseFiles(tpath+"indexpage.html", tpath+"navbar.html")
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "indexpage", data)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err.Error())
	}
}
