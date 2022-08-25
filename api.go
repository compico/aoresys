package main

import (
	"fmt"
	"net/http"

	"github.com/compico/aoresys/internal/userutil"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

func registerApiHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	model := r.FormValue("model")
	m := false
	if model == "alex" {
		m = true
	}
	if (model == "" || r.FormValue("username") == "") || (r.FormValue("email") == "" || r.FormValue("password") == "") {
		w.Write([]byte(fmt.Sprintf("Есть пустые поля!\n")))
		return
	}
	usr := model.User{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
		Model:    m,
	}
	v := userutil.NewValidator(usr.Username)
	if !v.ValidateByEqual() {
		w.Write([]byte(fmt.Sprintf("Недопустимые символы в никнейме\n")))
		return
	}
	if !v.ValidateLen() {
		w.Write([]byte(fmt.Sprintf("Недопустимая длинна в никнейме\n")))
		return
	}
	v = nil
	if len(usr.Password) < 8 || len(usr.Password) > 32 {
		w.Write([]byte(fmt.Sprintf("Недопустимая длинна в пароле\n")))
		return
	}
	hashpass, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Ошибка: %v\n", err)))
	}
	usr.Password = string(hashpass)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func loginApiHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}
