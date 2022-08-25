package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func Middleware(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
		// call registered handler
		n(w, r, ps)
	}

}
