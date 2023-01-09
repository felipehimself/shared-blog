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
	if _, err := statement.Exec(postId , userId); err != nil {
		return err
	}

	return nil

}
