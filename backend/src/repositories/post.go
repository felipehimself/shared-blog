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

	if _, err := statement.Exec(post.Title, post.Subtitle, post.AuthorId, post.Content); err != nil {
		return err
	}

	return nil

}

func (p *PostDB) GetPosts() ([]models.Post, error) {

	// TODO:
	// Aqui terá paginação... precisarei mudar a rota
	rows, err := p.db.Query("SELECT id, title, subtitle, author_id, votes_number, comments_number, ROUND((CHAR_LENGTH(content) / 200)) as minutesRead, created_at FROM posts")

	if err != nil {
		return []models.Post{}, err
	}

	var posts []models.Post

	for rows.Next() {

		var post models.Post

		if err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Subtitle,
			&post.AuthorId,
			&post.VotesNumber,
			&post.CommentsNumber,
			&post.MinutesRead,
			&post.CreatedAt); err != nil {
			return []models.Post{}, err
		}
		posts = append(posts, post)
	}

	return posts, nil

}
