package notes

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Notes struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

func (note Notes) Display() {
	fmt.Printf("Your note titled %v has the following content\n\n%v\n\n", note.Title, note.Content)

}

func (note Notes) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (Notes, error) {

	if title == "" || content == "" {
		return Notes{}, errors.New("invalid input")
	}

	return Notes{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil

}
