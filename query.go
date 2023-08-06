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
