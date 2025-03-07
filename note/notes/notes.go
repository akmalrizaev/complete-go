package notes

import (
	"errors"
	"time"
)

type Notes struct {
	title     string
	content   string
	createdAt time.Time
}

func New(title, content string) (Notes, error) {

	if title == "" || content == "" {
		return Notes{}, errors.New("invalid input")
	}

	return Notes{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil

}
