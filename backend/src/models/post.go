package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	Id             uint64    `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Subtitle       string    `json:"subtitle,omitempty"`
	Author         string    `json:"author,omitempty"`
	AuthorId       uint64    `json:"authorId,omitempty"`
	Content        string    `json:"content,omitempty"`
	VotesNumber    uint64    `json:"votesNumber,omitempty"`
	CommentsNumber uint64    `json:"commentsNumber,omitempty"`
	MinutesRead		 uint64			`json:"minutesRead,omitempty"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
}

// TODO:
// INICIALIZAR GIT MAS NA PASTA PRINCIPAL
// POSTAR IMAGEM?
// VERIFICAR SE O MODELS E A TABELA DE POST FAZEM SENTIDO OU SE PRECISA AJUSTAR
// CONTINUAR CONTROLLERS
// CRIAR REPOSITORIO DE POST

func (p *Post) ValidateFields() error {

	if p.Title == "" && p.Subtitle == "" && p.Content == "" {
		return errors.New("all fields are required")
	}

	if p.Title == "" {
		return errors.New("title is required")
	}

	if p.Subtitle == "" {
		return errors.New("subtitle is required")
	}

	if p.Content == "" {
		return errors.New("content is required")
	}

	p.cleanWhiteSpaces()

	return nil

}

func (p *Post) cleanWhiteSpaces() {

	p.Title = strings.TrimSpace(p.Title)
	p.Subtitle = strings.TrimSpace(p.Subtitle)
	p.Content = strings.TrimSpace(p.Content)

}
