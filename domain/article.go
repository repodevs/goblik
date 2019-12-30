package domain

import (
  "fmt"
  "errors"
  "strings"
  "time"
)


// Article errors
var (
  ErrCategoryNotFound = errors.New("category not found")
  ErrArticleNotFound = errors.New("article not found")
  ErrArticleExist = errors.New("article already exist")
)

// Article status
const (
  StatusArticlePublished = "Published"
  StatusArticleDraft = "Draft"
)

// Article categories
const (
  CategoryArticleInternet = "Internet"
  CategoryArticleProgramming = "Programming"
  CategoryArticlePhotography = "Photography"
)

// Article specification
type Article struct {
  Author Author
  Slug string
  Title string
  Body string
  Categories []string
  Status string
  CreatedAt int64
}

// ArticleHelper interface
type ArticleHelper interface {
  IsSlugExist(string) bool
}

// Generate slug e.g `this-is-title-timestamp`
func generateSlug(title string) string {
  arrTitle := strings.Split(title, "  ")
  suffix := fmt.Sprint(time.Now().UnixNano())
  suffix = suffix[len(suffix)/2:]
  arrTitle = append(arrTitle, suffix)
  // TODO: Lowercase the slug?
  return strings.Join(arrTitle, "-")
}

// CreateArticle create new article
func CreateArticle(helper ArticleHelper, slug, title, body string, author Author) (*Article, error) {
  // Check if article already exist
  if slug != "" && helper.IsSlugExist(slug) {
    return nil, ErrArticleExist
  }
  // Generate slug if not provided
  if slug == "" {
    slug = generateSlug(title)
  }
  // Return a valid article data
  return &Article{
    Slug: slug,
    Title: title,
    Body: body,
    Author: author,
    Status: StatusArticleDraft,
    CreatedAt: time.Now().UnixNano(),
  }, nil
}

// AddCategory to Article
func (a *Article) AddCategory(category string) error {
  switch category {
    case CategoryArticleInternet:
    case CategoryArticleProgramming:
    case CategoryArticlePhotography:
    default:
      return ErrCategoryNotFound
  }
  a.Categories = append(a.Categories, category)
  return nil
}

