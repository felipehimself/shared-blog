package repositories

import (
	"database/sql"
	"errors"
	"shared-blog-backend/src/models"

	"github.com/gofiber/fiber/v2"
)

type PostDB struct {
	db *sql.DB
}

func PostRepository(db *sql.DB) *PostDB {
	return &PostDB{db}
}

func (p *PostDB) Create(post models.Post) error {

	ts, err := p.db.Begin()

	if err != nil {
		return err
	}

	res, err := ts.Exec("INSERT INTO posts (title, subtitle, author_id, content) VALUES (?,?,?,?)",
		post.Title, post.Subtitle, post.AuthorId, post.Content)

	if err != nil {
		ts.Rollback()
		return err
	}

	lasId, err := res.LastInsertId()

	if err != nil {
		return err
	}

	if _, err := ts.Exec("INSERT INTO post_topics (topic_id, post_id) VALUES (?, ?)", post.TopicId, lasId); err != nil {
		ts.Rollback()
		return err
	}

	if err = ts.Commit(); err != nil {
		return err
	}

	return nil

}

func (p *PostDB) GetPosts(userId uint64) ([]models.Post, error) {

	rows, err := p.db.Query(`
	SELECT 
    p.id,
    u.username,
		u.name,
    t.topic,
    p.title,
    p.subtitle,
    p.author_id,
    COUNT(pc.comment) AS comments,
    COUNT(pv.post_id) AS votes,
			CASE WHEN (SELECT post_id from post_votes pvts WHERE pvts.user_id = ? AND pvts.post_id = pv.post_id)
				THEN 'true'
						ELSE 'false'
			END AS voted,
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
        LEFT JOIN
    post_comments pc ON pc.post_id = p.id
        LEFT JOIN
    post_votes pv ON p.id = pv.post_id
GROUP BY p.id

	`, userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []models.Post

	for rows.Next() {

		var post models.Post

		if err := rows.Scan(
			&post.Id,
			&post.Username,
			&post.Author,
			&post.Topic,
			&post.Title,
			&post.Subtitle,
			&post.AuthorId,
			&post.Comments,
			&post.Votes,
			&post.Voted,
			&post.MinutesRead,
			&post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil

}

func (p *PostDB) Vote(postId, userId uint64) error {

	statement, err := p.db.Prepare("INSERT INTO post_votes (post_id, user_id) VALUES (?,?)")

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(postId, userId); err != nil {
		return err
	}

	return nil
}

func (p *PostDB) UnVote(postId, userId uint64) error {

	statement, err := p.db.Prepare(`DELETE from post_votes WHERE post_id = ? AND user_id = ?`)

	if err != nil {
		return err
	}

	defer statement.Close()

	res, err := statement.Exec(postId, userId)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("this post doesn't exist")
	}

	return nil

}

func (p *PostDB) GetPost(postId, userId uint64) (models.Post, error) {

	row := p.db.QueryRow(`
	SELECT 
    p.id,
    u.username,
		u.name,
    t.topic,
    p.title,
    p.subtitle,
		p.content,
    p.author_id,
    COUNT(pc.comment) AS comments,
    COUNT(pv.post_id) AS votes,
			CASE WHEN (SELECT post_id from post_votes pvts WHERE pvts.user_id = ? AND pvts.post_id = pv.post_id)
				THEN 'true'
						ELSE 'false'
			END AS voted,
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
        LEFT JOIN
    post_comments pc ON pc.post_id = p.id
        LEFT JOIN
    post_votes pv ON p.id = pv.post_id
		WHERE p.id = ?
GROUP BY p.id
	
	`, userId, postId)

	var post models.Post

	if err := row.Scan(
		&post.Id,
		&post.Username,
		&post.Author,
		&post.Topic,
		&post.Title,
		&post.Subtitle,
		&post.Content,
		&post.AuthorId,
		&post.Comments,
		&post.Votes,
		&post.Voted,
		&post.MinutesRead,
		&post.CreatedAt); err != nil {
		return models.Post{}, err

	}

	return post, nil
}

func (p *PostDB) GetUserPosts(username string) ([]models.Post, error) {

	rows, err := p.db.Query(`
	SELECT 
		p.id as post_id, p.title, p.subtitle, p.created_at, u.name, u.username
	FROM
		posts p
			INNER JOIN
				users u ON u.id = p.author_id
	WHERE u.username = ?
	`, username)

	if err != nil {
		return nil, err
	}

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		if err = rows.Scan(&post.Id, &post.Title, &post.Subtitle, &post.CreatedAt, &post.Author, &post.Username); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil

}

func (p *PostDB) EditPost(postId, userOnToken uint64, post models.Post) (int, error) {

	row := p.db.QueryRow("SELECT author_id FROM posts WHERE id = ?", postId)

	var author models.User

	if err := row.Scan(&author.ID); err != nil {
		return fiber.StatusBadRequest, err
	}

	if author.ID != userOnToken {
		return fiber.StatusUnauthorized, errors.New("you cannot edit a post from another user")
	}

	ts, err := p.db.Begin()

	if err != nil {
		ts.Rollback()
		return fiber.StatusInternalServerError, err
	}

	statement, err := ts.Prepare(`UPDATE posts SET title = ?, subtitle = ?, content = ?  WHERE id = ?`)

	if err != nil {
		ts.Rollback()
		return fiber.StatusInternalServerError, err
	}

	if _, err := statement.Exec(post.Title, post.Subtitle, post.Content, postId); err != nil {
		ts.Rollback()
		return fiber.StatusInternalServerError, err
	}

	statement, err = ts.Prepare("UPDATE post_topics SET topic_id = ? WHERE post_id = ?")

	if err != nil {
		ts.Rollback()
		return fiber.StatusInternalServerError, err
	}

	if _, err := statement.Exec(post.TopicId, postId); err != nil {
		ts.Rollback()
		return fiber.StatusInternalServerError, err
	}

	if err = ts.Commit(); err != nil {
		return fiber.StatusInternalServerError, nil
	}

	return fiber.StatusOK, nil

}
