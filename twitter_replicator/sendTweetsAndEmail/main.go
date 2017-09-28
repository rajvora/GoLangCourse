// This is project is show in udemy course by Todd McLeod
// To run use the google application engine: https://cloud.google.com/sdk/docs/


package main

import (
	"encoding/json"
	"github.com/dustin/go-humanize"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	router := httprouter.New()
	http.Handle("/", router)

	// client page
	router.GET("/", homePage)
	router.GET("/user/:user", userPage)
	router.GET("/form/login", loginPage)
	router.GET("/form/signup", signupPage)

	// client api
	router.POST("/api/checkusername", checkUserName)
	router.POST("/api/createuser", createUser)
	router.POST("/api/login", loginProcess)
	router.POST("/api/tweet", tweetProcess)
	router.GET("/api/logout", logout)

	// example url http://localhost:8080/api/follow/<username>
	router.GET("/api/follow/:user", follow)
	router.GET("/api/unfollow/:user", unfollow)

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))

	tpl = template.New("roottemplate")
	tpl = tpl.Funcs(template.FuncMap{
		"humanize_time": humanize.Time,
	})
	tpl = template.Must(tpl.ParseGlob("templates/html/*.html"))

	appengine.Main()
}

func homePage(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	ctx := appengine.NewContext(req)
	//get tweets
	tweets, err := getTweets(req, nil)
	if err != nil {
		log.Errorf(ctx, "error getting tweets: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}
	// get session
	memItem, err := getCachedUser(req)
	var sd SessionData
	if err == nil {
		// logged in
		json.Unmarshal(memItem.Value, &sd)
		sd.LoggedIn = true
	}
	sd.Tweets = tweets
	tpl.ExecuteTemplate(res, "home.html", &sd)
}

func userPage(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	ctx := appengine.NewContext(req)
	user := User{UserName: ps.ByName("user")}
	//get tweets
	tweets, err := getTweets(req, &user)
	if err != nil {
		log.Errorf(ctx, "error getting tweets: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}
	// get session
	memItem, err := getCachedUser(req)
	var sd SessionData
	if err == nil {
		// logged in
		json.Unmarshal(memItem.Value, &sd)
		sd.LoggedIn = true
		sd.ViewingUser = user.UserName
		sd.FollowingUser, err = following(sd.UserName, user.UserName, req)
		if err != nil {
			log.Errorf(ctx, "error running following query: %v", err)
			http.Error(res, err.Error(), 500)
			return
		}
	}
	sd.Tweets = tweets
	tpl.ExecuteTemplate(res, "user.html", &sd)

}

func loginPage(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	serveTemplate(res, req, "login.html")
}

func signupPage(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	serveTemplate(res, req, "signup.html")
}
