package main

import "time"

type User struct {
	Email    string
	UserName string `datastore:"-"`
	Password string `json:"-"`
}

// data stored into the app engine datastor
type SessionData struct {
	User
	LoggedIn      bool
	LoginFail     bool
	Tweets        []Tweet
	ViewingUser   string
	FollowingUser bool
}

// store the message, time and userName
type Tweet struct {
	Msg      string
	Time     time.Time
	UserName string
}
