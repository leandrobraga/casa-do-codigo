package data

import (
	"regexp"
	"time"

	v "github.com/cohesivestack/valgo"
)

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type AuthorDtoInput struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

type Author struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"created_at"`
}

// func NewAuthor(i AuthorDtoInput) *Author {
// 	return &Author{
// 		Name: i.Name,

// 	}
// }

func (a Author) Validate() error {
	val := v.Is(
		v.String(a.Name, "name").Not().Blank(),
		v.String(a.Email, "email").Not().Blank(),
		v.String(a.Email, "email").Not().MatchingTo(EmailRX),
		v.String(a.Description, "description").Not().Empty().MaxLength(450),
		v.Time(a.CreateAt, "createAt").Not().Zero(),
	)

	if !val.Valid() {
		return val.Error()
	}
	return nil
}
