package service

import (
  "github.com/repodevs/goblik/domain"
  "github.com/repodevs/goblik/query"
)

// ArticleService implementation for article service
// this is for connectin between domain. repository and query
type ArticleService struct {
  Query query.ArticleQuery
}

// CreateArticleService create new article service
func CreateArticleService(query query.ArticleQuery) domain.ArticleHelper {
  return &ArticleService {
    Query: query,
  }
}

// Isslugexist is method for checking slug is exist or not
func (s *ArticleService) IsSlugExist(slug string) bool {
  _, err := s.Query.GetBySlug(slug)
  if err != nil {
    return false
  }
  return true
}
