package main

import (
  "fmt"

  "github.com/repodevs/goblik/repository"
  "github.com/repodevs/goblik/storage"
  "github.com/repodevs/goblik/query"
  "github.com/repodevs/goblik/domain"
  "github.com/repodevs/goblik/service"
)


func main() {
  // init storage
  inStorage := storage.CreateArticleStorage()
  q := query.CreateArticleQueryInStorage(inStorage)
  authorRepo := repository.CreateAuthorRepositoryInStorage(inStorage)

  articleRepo := repository.CreateArticleRepositoryInStorage(inStorage)
  s := service.CreateArticleService(q)

  // create an author
  authorOne, err := domain.CreateAuthor("Author One")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(*authorOne)
  err = authorRepo.Create(authorOne)
  if err != nil {
    fmt.Println(err)
  }

  // create article
  articleOne, err := domain.CreateArticle(s, "", "this is article", "body", *authorOne)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println("[+] Article 1 -> ", *articleOne)

  // assign categories
  articleOne.AddCategory(domain.CategoryArticleInternet)
  articleOne.AddCategory(domain.CategoryArticlePhotography)
  articleRepo.Create(articleOne)
  fmt.Println(articleOne)
  fmt.Println()

  // create article 2
  articleTwo, err := domain.CreateArticle(s, "", "this is article two", "body article two", *authorOne)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println("[+] Article 2 ->", *articleTwo)

  // assign category to article two
  articleTwo.AddCategory(domain.CategoryArticleProgramming)
  articleRepo.Create(articleTwo)
  fmt.Println(articleTwo)
  fmt.Println()

  // add article to pages
  authorOne.AddPage(*articleOne)
  authorOne.AddPage(*articleTwo)
  authorRepo.Update(authorOne.ID, *authorOne)
  fmt.Println(authorOne)
  fmt.Println()

  // print all article in storage
  fmt.Println(inStorage.Article)
  fmt.Println()
  fmt.Println(inStorage.Author)
  fmt.Println()

}

