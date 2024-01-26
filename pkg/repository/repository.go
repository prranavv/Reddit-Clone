package repository

import "github.com/prranavv/reddit_clone/pkg/models"

type DBrepo interface {
	InsertUserDetails(username, email, password string) error
	GetPasswordFromUsername(username string) (string, error)
	GetUserIDFromUsername(username string) (int, error)
	CreatePost(Username, title, body, subreddit, image_path, video_path string) error
	GetingPostsFromSubreddit(subreddit string) ([]models.Post, error)
	InsertingDataIntoLikedTable(post_id int) error
	GettingPostIDFromDetails(username, title, body, subreddit string) (int, error)
	DeletePost(post_id int) error
	CheckingDuplicatePost(body, title, subreddit, username string) ([]int, error)
	GettingLikes(post_id int) (int, error)
	AddingLikes(post_id int) error
	Disliking(post_id int) error
}
