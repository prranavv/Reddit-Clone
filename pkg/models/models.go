package models

import "time"

type User struct {
	User_ID  int
	Email    string
	Username string
	Password string
}

type Post struct {
	Post_ID    int
	Username   string
	Title      string
	Body       string
	Liked      Liked
	Created_at time.Time
	Updated_at time.Time
	ImageUrl   string
}

type Liked struct {
	Post_ID    int
	Liked      bool
	Likes      int
	Created_at time.Time
	Updated_at time.Time
}
