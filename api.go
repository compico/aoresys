package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"net/mail"
	"strings"
	"time"

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
	hashpass, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Ошибка: %v\n", err)))
	}
	usr.Password = string(hashpass)
	err = cdb.AddNewUser(ctx, usr)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Ошибка: %v\n", err)))
		return
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func registrationHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	mn := r.FormValue("model")
	var data = struct {
		Errors        []string
		User          userutil.User
		UsernameError bool
		EmailError    bool
		PasswordError bool
	}{
		User: userutil.User{
			Username: r.FormValue("username"),
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		},
	}
	var m bool
	if mn == "" {
		data.Errors = append(data.Errors, "Model not selected!")
	}
	if mn == "alex" {
		m = userutil.ALEX
	}
	data.User.Model = m
	t, err := template.ParseFiles(tpath + "registationForm.html")
	if err != nil {
		data.Errors = append(data.Errors, err.Error())
	}

	//email
	if data.User.Email == "" {
		data.EmailError = true
		data.Errors = append(data.Errors, "Email is Empty!")
	}
	if _, err := mail.ParseAddress(data.User.Email); err != nil {
		data.EmailError = true
		data.Errors = append(data.Errors, "Email is not valid!")
	}

	//username
	if data.User.Username == "" {
		data.UsernameError = true
		data.Errors = append(data.Errors, "Username is Empty!")
	}
	v := userutil.NewValidator(data.User.Username)
	if v.ValidateByEqual() {
		data.UsernameError = true
		data.Errors = append(data.Errors, "Username has wrong characters!")
	}
	if v.ValidateLen() {
		data.UsernameError = true
		data.Errors = append(data.Errors, "Invalid username length!")
	}
	v = nil

	//password
	if data.User.Password == "" {
		data.PasswordError = true
		data.Errors = append(data.Errors, "Password is Empty!")
	}
	if len(data.User.Password) < 4 || len(data.User.Password) > 32 {
		data.PasswordError = true
		data.Errors = append(data.Errors, "Invalid password length!")
	}

	t.ExecuteTemplate(w, "registrationform", data)
}

func loginApiHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}
