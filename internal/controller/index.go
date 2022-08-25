package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpath = "./templates/"

func IndexController(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
