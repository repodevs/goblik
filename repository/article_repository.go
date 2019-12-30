package repository

import (
  "github.com/repodevs/goblik/storage"
  "github.com/repodevs/goblik/domain"
)

type ArticleRepository interface {
  Create(*domain.Article) error
}

//ArticleRepositoryInStorage implementation
type ArticleRepositoryInStorage struct {
  Storage *storage.ArticleStorage
}

// CreateArticleRepositoryInStorage instantiate new article repository in storage
func CreateArticleRepositoryInStorage(storage *storage.ArticleStorage) ArticleRepository {
  return &ArticleRepositoryInStorage {
    Storage: storage,
  }
}

// Create new article
func (r *ArticleRepositoryInStorage) Create(article *domain.Article) error {
  r.Storage.Article[article.Slug] = *article
  return nil
}

