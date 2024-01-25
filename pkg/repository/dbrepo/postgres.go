package dbrepo

import (
	"context"
	"database/sql"
	"time"

	"github.com/prranavv/reddit_clone/pkg/models"
)

func (m *postgresRepo) InsertUserDetails(username, email, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `INSERT INTO public."user"(
		email, password, username)
		VALUES ($1,$2,$3);`
	_, err := m.DB.ExecContext(ctx, query, email, password, username)
	if err != nil {
		return err
	}
	return nil
}

func (m *postgresRepo) GetPasswordFromUsername(username string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var name string
	query := `SELECT password
	FROM public."user" where username=$1;`
	err := m.DB.QueryRowContext(ctx, query, username).Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

func (m *postgresRepo) GetUserIDFromUsername(username string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var id int
	query := `SELECT user_id
	FROM public."user" where username=$1;`
	err := m.DB.QueryRowContext(ctx, query, username).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (m *postgresRepo) CreatePost(username, title, body, subreddit, image_path string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `INSERT INTO public.post(
		username, title, body, created_at, updated_at, subreddit,image_path)
		VALUES ($1,$2,$3,$4,$5,$6,$7);`
	_, err := m.DB.ExecContext(ctx, query,
		username,
		title,
		body,
		time.Now(),
		time.Now(),
		subreddit,
		image_path,
	)
	if err != nil {
		return err
	}
	return nil
}

func (m *postgresRepo) GetingPostsFromSubreddit(subreddit string) ([]models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var Posts []models.Post
	query := `SELECT p.post_id, p.title, p.body,p.username,p.image_path,l.no_of_likes
	FROM public.post p join liked l on p.post_id=l.post_id where subreddit=$1 order by post_id desc	;`
	rows, err := m.DB.QueryContext(ctx, query, subreddit)
	if err != nil {
		return Posts, err
	}
	var imgURL sql.NullString
	defer rows.Close()
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Post_ID,
			&post.Title,
			&post.Body,
			&post.Username,
			&imgURL,
			&post.Liked.Likes,
		)
		if imgURL.Valid {
			post.ImageUrl = imgURL.String
		}
		if err != nil {
			return Posts, err
		}
		Posts = append(Posts, post)
	}
	if err := rows.Err(); err != nil {
		return Posts, err
	}
	return Posts, nil
}

func (m *postgresRepo) InsertingDataIntoLikedTable(post_id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `INSERT INTO public.liked(
		post_id)
		VALUES ($1);`
	_, err := m.DB.ExecContext(ctx, query, post_id)
	if err != nil {
		return err
	}
	return nil
}

func (m *postgresRepo) GettingPostIDFromDetails(username, title, body, subreddit string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var post_id int
	query := `SELECT post_id
	FROM public.post where username=$1 and title=$2 and body=$3 and subreddit=$4;`
	err := m.DB.QueryRowContext(ctx, query, username, title, body, subreddit).Scan(&post_id)
	if err != nil {
		return 0, err
	}
	return post_id, nil
}

func (m *postgresRepo) DeletePost(post_id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `DELETE FROM public.post
	WHERE post_id=$1;`
	_, err := m.DB.ExecContext(ctx, query, post_id)
	if err != nil {
		return err
	}
	return nil
}

func (m *postgresRepo) CheckingDuplicatePost(body, title, subreddit, username string) ([]int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `SELECT post_id
	FROM public.post where username=$1 and subreddit=$2 and title=$3 and body=$4;`
	var post_ids []int
	rows, err := m.DB.QueryContext(ctx, query, username, subreddit, title, body)
	if err != nil {
		return post_ids, err
	}
	defer rows.Close()
	for rows.Next() {
		var post_id int
		err := rows.Scan(
			&post_id,
		)
		if err != nil {
			return post_ids, err
		}
		post_ids = append(post_ids, post_id)
	}
	if err := rows.Err(); err != nil {
		return post_ids, err
	}
	return post_ids, nil
}

func (m *postgresRepo) GettingLikes(post_id int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `SELECT no_of_likes
	FROM public.liked where post_id=$1;`
	var no_of_likes int
	row := m.DB.QueryRowContext(ctx, query, post_id)
	err := row.Scan(&no_of_likes)
	if err != nil {
		return no_of_likes, err
	}
	return no_of_likes, nil
}

func (m *postgresRepo) AddingLikes(post_id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	no_of_likes, err := m.GettingLikes(post_id)
	if err != nil {
		return err
	}
	number_to_be_added := no_of_likes + 1
	query := `UPDATE public.liked
	SET no_of_likes=$1
	WHERE post_id=$2;`
	_, err = m.DB.ExecContext(ctx, query, number_to_be_added, post_id)
	if err != nil {
		return err
	}
	return nil
}

func (m *postgresRepo) Disliking(post_id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	no_of_likes, err := m.GettingLikes(post_id)
	if err != nil {
		return err
	}
	number_to_be_added := no_of_likes - 1
	query := `UPDATE public.liked
	SET no_of_likes=$1
	WHERE post_id=$2;`
	_, err = m.DB.ExecContext(ctx, query, number_to_be_added, post_id)
	if err != nil {
		return err
	}
	return nil
}
