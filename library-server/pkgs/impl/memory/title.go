package memory

import (
	"context"
	"libstack/pkgs/model"

	"github.com/pkg/errors"
)

type TitleReadWriter struct {
	cache map[string]model.Title
}

func NewTitleReadWriter() *TitleReadWriter {
	return &TitleReadWriter{cache: map[string]model.Title{}}
}

func (l *TitleReadWriter) GetByIsbn(ctx context.Context, isbn string) (title model.Title, err error) {
	if isbn == "" {
		err = errors.New("isbn required")
		return
	}
	title, ok := l.cache[isbn]
	if !ok {
		err = errors.Errorf("could not find title with isbn %q", isbn)
		return
	}
	return title, nil
}

// TODO(mchenryc): should I be returning a pointer here? That seems like it would mess up the cache

func (l *TitleReadWriter) Add(ctx context.Context, title model.Title) (model.Title, error) {
	// TODO(mchenryc): logging
	// TODO(mchenryc): should validation happen somewhere else?
	empty := model.Title{}
	if title.Isbn == "" {
		return empty, errors.New("isbn required")
	}
	if _, ok := l.cache[title.Isbn]; ok {
		return empty, errors.New("title with isbn already exists")
	}
	l.cache[title.Isbn] = title
	return title, nil
}
