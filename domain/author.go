package domain

import (
  "errors"
  "time"
)

// MaxAuthorPages is max pages that can be created by author
const MaxAuthorPages = 3

// Errors specification
var (
  ErrAuthorNotFound = errors.New("author not found")
  ErrExceededPages = errors.New("exceeded pages")
)

// Author is user who create an article
type Author struct {
  ID string
  Pages []Article
  Name string
  CreatedAt int64
}

// CreateAuthor create a new author
func CreateAuthor(name string) (*Author, error) {
  return &Author{
    Name: name,
    CreatedAt: time.Now().UnixNano(),
  }, nil
}

// AddPage to author pages
func (a *Author) AddPage(article Article) error {
  if len(a.Pages)+1 > MaxAuthorPages {
    return ErrExceededPages
  }
  a.Pages = append(a.Pages, article)
  return nil
}

