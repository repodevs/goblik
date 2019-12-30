package query

import (
  "github.com/repodevs/goblik/storage"
  "github.com/repodevs/goblik/domain"
)


// ArticleQuery interface
type ArticleQuery interface {
  GetBySlug(string) (domain.Article, error)
}

type ArticleQueryInStorage struct {
  Storage *storage.ArticleStorage
}

// CreateArticleQueryInStorage instantiate new article query in storage
func CreateArticleQueryInStorage(storage *storage.ArticleStorage) ArticleQuery {
  return &ArticleQueryInStorage {
    Storage: storage,
  }
}

// GetBySlug in article storage
func (q *ArticleQueryInStorage) GetBySlug(slug string) (domain.Article, error) {
  article, ok := q.Storage.Article[slug]
  if !ok {
    return article, domain.ErrArticleNotFound
  } else {
    return article, nil
  }
}

