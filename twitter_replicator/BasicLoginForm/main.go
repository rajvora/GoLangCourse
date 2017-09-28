package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

// To run us

var tpl *template.Template

func main() {
	r := httprouter.New()
	http.Handle("/", r)
	r.GET("/", Home)
	r.GET("/login", Login)
	r.GET("/signup", Signup)

	// access public files
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	tpl = template.Must(template.ParseGlob("templates/html/*.html"))

	http.ListenAndServe(":8080",nil)
}

func Home(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	tpl.ExecuteTemplate(res, "home.html", nil)
}

func Login(res http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(res, "login.html", nil)
}

func Signup(res http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(res, "signup.html", nil)
}
