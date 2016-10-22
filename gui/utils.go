package gui

import (
	"net/http"

	"github.com/skycoin/skyhash/skyhashmanager"
)

type SkyhashManager struct {
	*skyhashmanager.SkyhashManager
}

// GET returns StatusMethodNotAllowed if the method is not GET
func GET(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}

// POST returns StatusMethodNotAllowed if the method is not POST
func POST(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handler(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}

// DELETE returns StatusMethodNotAllowed if the method is not DELETE
func DELETE(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			handler(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}

type mRoute struct {
	method  string
	handler func(w http.ResponseWriter, r *http.Request)
}

func MethodToHandler(method string, handler func(w http.ResponseWriter, r *http.Request)) *mRoute {
	return &mRoute{
		method:  method,
		handler: handler,
	}
}

// MethodMux selects a mRoute based on the method of the request
func MethodMux(methodMuxes ...*mRoute) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, m := range methodMuxes {
			if r.Method == m.method {
				m.handler(w, r)
				return
			}
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
