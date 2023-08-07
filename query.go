package main

func GetAllPosts(offset int) ([]Post, error) {
	var posts = []Post{}
	rows, err := db.Query(`
		SELECT id, title, slug, body, published
		FROM posts
		ORDER BY published
		LIMIT $1 OFFSET $2`, PageLimit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Title, &post.Slug, &post.Body, &post.Published)
		if err == nil {
			posts = append(posts, post)
		} else {
			return nil, err
		}
	}
	return posts, nil
}

func GetSinglePost(id int, slug string) (Post, error) {
	var post Post
	var err error
	row := db.QueryRow(`
		SELECT id, title, slug, body, published
		FROM posts
		WHERE id = $1 AND slug = $2`, id, slug)
	err = row.Scan(&post.Id, &post.Title, &post.Slug, &post.Body, &post.Published)
	return post, err
}

func CountPosts() (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM posts").Scan(&count)
	if err == nil {
		return count, nil
	}
	return count, err
}

func GetPostById(id int) (*Post, error) {
	var post Post
	var err error
	row := db.QueryRow(`
		SELECT id, title, slug, body, published
		FROM posts
		WHERE id = $1`, id)
	err = row.Scan(&post.Id, &post.Title, &post.Slug, &post.Body, &post.Published)
	return &post, err
}

func CreateComment(postid int, name, body string) error {
	var err error
	result, err := db.Exec(`INSERT INTO comments (post, name, body) VALUES ($1, $2, $3)`, postid, name, body)
	if err == nil {
		// LastInsertId not supported by this postgres driver
		_, err = result.RowsAffected()
	}
	return err
}

func AllPostComments(postid int) ([]Comment, error) {
	var comments []Comment
	rows, err := db.Query(`SELECT id, name, body, created FROM comments WHERE post = $1`, postid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		comment := Comment{}
		err = rows.Scan(&comment.Id, &comment.Name, &comment.Body, &comment.Created)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
