package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/compico/aoresys/aoreblg"
	"github.com/julienschmidt/httprouter"
)

var tpath = "./templates/"
var x = aoreblg.Posts{}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles(
		tpath+"head.html",
		tpath+"index.html",
		tpath+"topmenu.html",
	)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "index", x)
	if err != nil {
		fmt.Fprintf(w, "[ERROR] %v!!", err.Error())
		fmt.Printf("[ERROR] %v!!", err.Error())
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles(
		tpath+"head.html",
		tpath+"register.html",
		tpath+"topmenu.html",
	)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "register", nil)
	if err != nil {
		fmt.Fprintf(w, "[ERROR] %v!!", err.Error())
		fmt.Printf("[ERROR] %v!!", err.Error())
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles(
		tpath+"head.html",
		tpath+"login.html",
		tpath+"topmenu.html",
	)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err.Error())
	}
	err = t.ExecuteTemplate(w, "login", nil)
	if err != nil {
		fmt.Fprintf(w, "[ERROR] %v!!", err.Error())
		fmt.Printf("[ERROR] %v!!", err.Error())
	}
}

func testdata() {
	x.Posts = make([]aoreblg.Post, 0)

	for i := 0; i < 5; i++ {
		y := aoreblg.Post{
			Title:        "Проверка названия блога " + strconv.Itoa(i+1),
			Date:         time.Now().Round(2 * time.Second),
			PreviewImage: "/public/images/animals.png",
			PreviewText:  "Test проверка текста или прочего дерьма  Test проверка текста или прочего дерьмаTest проверка текста или прочего дерьма  Test проверка текста или прочего дерьмаTest проверка текста или прочего дерьма  Test проверка текста или прочего дерьмаTest проверка текста или прочего дерьма  Test проверка текста или прочего дерьма",
			Text:         "HELLLLOOWWW EvRyVaNjeiawoejioawjoie",
		}
		x.Posts = append(x.Posts, y)
	}
}
