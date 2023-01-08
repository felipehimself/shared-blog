package repositories

import (
	"database/sql"
	"shared-blog-backend/src/models"
)

type PostDB struct {
	db *sql.DB
}

func PostRepository(db *sql.DB) *PostDB {
	return &PostDB{db}
}

func (p *PostDB) Create(post models.Post) error {

	statement, err := p.db.Prepare("INSERT INTO posts (title, subtitle, author_id, content) VALUES (?,?,?,?)")

	if err != nil {
		return err
	}

	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Subtitle, post.AuthorId, post.Content)

	if err != nil {
		return err
	}

	lasId, err := result.LastInsertId()

	if err != nil {
		return err
	}

	statement, err = p.db.Prepare("INSERT INTO post_topics (topic_id, post_id) VALUES (?, ?)")

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(post.TopicId, lasId); err != nil {
		return err
	}

	return nil

}

func (p *PostDB) GetPosts() ([]models.Post, error) {

	rows, err := p.db.Query(`
	SELECT 
		p.id,
		u.username,
		t.topic,
		p.title,
		p.subtitle,
		p.author_id,
		p.votes,
		p.comments,
		ROUND((CHAR_LENGTH(p.content) / 200)) AS minutes_read,
		p.created_at
	FROM
		posts p
				INNER JOIN
		users u ON p.author_id = u.id
				INNER JOIN
		post_topics pt ON p.id = pt.post_id
				INNER JOIN
		topics t ON t.id = pt.topic_id
	`)

	if err != nil {
		return nil, err
	}

	var posts []models.Post

	for rows.Next() {

		var post models.Post

		if err := rows.Scan(
			&post.Id,
			&post.Username,
			&post.Topic,
			&post.Title,
			&post.Subtitle,
			&post.AuthorId,
			&post.Votes,
			&post.Comments,
			&post.MinutesRead,
			&post.CreatedAt); err != nil {
			return []models.Post{}, err
		}
		posts = append(posts, post)
	}

	return posts, nil

}

func (p *PostDB) Vote(postId uint64) error {

	statement, err := p.db.Prepare("UPDATE posts SET votes = votes + 1 WHERE id = ?")

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(postId); err != nil {
		return err
	}

	return nil
}

func (p *PostDB) UnVote(postId uint64) error {

	statement, err := p.db.Prepare(`
	UPDATE posts 
	SET 
    votes = CASE
        WHEN votes - 1 < 0 THEN 0
        ELSE votes - 1 
    END
	WHERE
    posts.id = ?
	`)

	if err != nil {
		return err
	}

	defer statement.Close()
	if _, err := statement.Exec(postId); err != nil {
		return err
	}

	return nil

}
