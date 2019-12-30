package storage

import (
  "github.com/repodevs/goblik/domain"
)

// ArticleStorage used to temporary store article's data in memory
type ArticleStorage struct {
  Article map[string]domain.Article
  Author map[string]domain.Author
}

// CreateArticleStorage create new storage
func CreateArticleStorage() *ArticleStorage {
  return &ArticleStorage {
    Article: make(map[string]domain.Article),
    Author: make(map[string]domain.Author),
  }
}

