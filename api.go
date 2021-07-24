package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/compico/aoresys/internal/userutil"
	"github.com/julienschmidt/httprouter"
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
	usr := userutil.User{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
		Model:    m,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	b, err := cdb.ExistUsername(ctx, strings.ToLower(usr.Username))
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Ошибка: %v\n", err)))
		return
	}
	if b {
		w.Write([]byte("Никнейм занят!"))
		return
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
	err = cdb.AddNewUser(ctx, usr)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Ошибка: %v\n", err)))
		return
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
