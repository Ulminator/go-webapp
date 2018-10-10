package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"../viewmodel"
)

type home struct {
	homeTemplate         *template.Template
	standLocatorTemplate *template.Template
	loginTemplate        *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/login", h.handleLogin)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewHome()
	// Browser can interpret this, except when compressed
	w.Header().Add("Content-Type", "text/html")
	// time.Sleep(4 * time.Second)
	h.homeTemplate.Execute(w, vm)
}

func (h home) handleLogin(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewLogin()

	if r.Method == http.MethodPost {
		err := r.ParseForm() // ends up populating r.Form map
		if err != nil {
			log.Println(fmt.Errorf("Error logging in: %v", err))
		}
		// r.Form is a map of kv pairs
		email := r.Form.Get("email")
		password := r.Form.Get("password")
		if email == "test@gmail.com" && password == "password" {
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
			return
		} else {
			// Populates the form on failed login attempt.
			vm.Email = email
			vm.Password = password
		}
	}
	w.Header().Add("Content-Type", "text/html")
	h.loginTemplate.Execute(w, vm)
}
