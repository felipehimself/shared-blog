package repositories

import (
	"database/sql"
	"errors"
	"shared-blog-backend/src/models"

	"github.com/gofiber/fiber/v2"
)

type CommentDB struct {
	db *sql.DB
}

func CommentRepository(db *sql.DB) *CommentDB {
	return &CommentDB{db}
}

func (c *CommentDB) CommentPost(userOnToken uint64, comment models.Comment) error {

	statement, err := c.db.Prepare("INSERT INTO post_comments (user_id, post_id, comment) VALUES (?,?,?)")

	if err != nil {
		return err
	}

	if _, err = statement.Exec(userOnToken, comment.PostId, comment.Comment); err != nil {
		return err
	}

	return nil

}

func (c *CommentDB) DeleteComment(userOnToken, commentId uint64) (int, error) {

	row := c.db.QueryRow("SELECT pc.user_id FROM post_comments pc WHERE pc.id = ?", commentId)

	var userId uint64

	if err := row.Scan(&userId); err != nil {
		return fiber.StatusBadRequest, err
	}

	if userId != userOnToken {
		return fiber.StatusForbidden, errors.New("you cannot delete a post form other user")
	}

	statement, err := c.db.Prepare("DELETE from post_comments WHERE id = ?")

	if err != nil {
		return fiber.StatusInternalServerError, err
	}

	if _, err := statement.Exec(commentId); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil

}
