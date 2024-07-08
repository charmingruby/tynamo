package factory

import (
	"time"

	"github.com/charmingruby/tynamo/internal/model"
	"github.com/oklog/ulid/v2"
)

func MakeExample(title string) *model.Example {
	id := ulid.Make().String()

	return &model.Example{
		ID:        id,
		Title:     title,
		CreatedAt: time.Now(),
	}
}
