package main

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"net/http"
)

// used the app engine to keep track of the session
func getCachedUser(req *http.Request) (*memcache.Item, error) {
	ctx := appengine.NewContext(req)

	cookie, err := req.Cookie("session")
	if err != nil {
		return &memcache.Item{}, err
	}

	item, err := memcache.Get(ctx, cookie.Value)
	if err != nil {
		return &memcache.Item{}, err
	}
	return item, nil
}
