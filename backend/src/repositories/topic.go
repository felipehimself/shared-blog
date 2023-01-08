package repositories

import (
	"database/sql"
	"shared-blog-backend/src/models"
)

type TopicDB struct {
	db *sql.DB
}

func TopicRepository(db *sql.DB) *TopicDB {
	return &TopicDB{db}
}

func (t *TopicDB) Topics() ([]models.Topic, error) {

	var topics []models.Topic

	rows, err := t.db.Query("SELECT * FROM topics")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var topic models.Topic

		if err := rows.Scan(&topic.Id, &topic.Topic, &topic.CreatedAt); err != nil {
			return nil, err
		}

		topics = append(topics, topic)

	}

	return topics, nil

}
