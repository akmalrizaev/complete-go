package notes

import (
	"errors"
	"fmt"
	"time"
)

type Notes struct {
	title     string
	content   string
	createdAt time.Time
}

func (note Notes) Display() {
	fmt.Printf("Your note titled %v has the following content\n\n%v\n\n", note.title, note.content)

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
