package repository

import (
  "fmt"
  "time"

  "github.com/repodevs/goblik/domain"
  "github.com/repodevs/goblik/storage"
)

type AuthorRepository interface {
  Create(*domain.Author) error
  Update(string, domain.Author) error
}

type AuthorRepositoryInStorage struct {
  Storage *storage.ArticleStorage
}

// CreateAuthorRepositoryInStorage instance
func CreateAuthorRepositoryInStorage(storage *storage.ArticleStorage) AuthorRepository {
  return &AuthorRepositoryInStorage{
    Storage: storage,
  }
}

// Update author based on id
func (r *AuthorRepositoryInStorage) Update(slug string, author domain.Author) error {
  if _, ok := r.Storage.Author[slug]; !ok {
    return domain.ErrAuthorNotFound
  }
  r.Storage.Author[slug] = author
  return nil
}

// Create new author
func (r *AuthorRepositoryInStorage) Create(author *domain.Author) error {
  author.ID = fmt.Sprint(time.Now().UnixNano())
  r.Storage.Author[author.ID] = *author
  return nil
}
