// ArtcleRepository_test.go
package repository

import (
	"blogo/go/domain"
	"fmt"
	"testing"
)

func TestGetOneArticle(t *testing.T) {
	aritlcleSave := domain.Article{}
	aritlcleSave.Subject = "안녕?"
	aritlcleSave.Content = "헬로월드"
	aritlcleSave.Usernick = "아라한사"
	id, err := Save(&aritlcleSave)
	checkErr(err)
	article := GetOneArticle(id)
	fmt.Println(article)
	fmt.Println("수정시도")
	article.Subject = "수정된 게시글"
	Update(&article)
	article = GetOneArticle(id)
	fmt.Println(article)
}

func TestDeleteArticle(t *testing.T) {
	DeleteArticle(25)
	article := GetOneArticle(25)
	fmt.Println(article)
}

func TestGetArticleList(t *testing.T) {
	fmt.Println("===== 게시글들 =====")
	articles := GetArticleList()
	fmt.Println(articles)
}

func TestCheckErr(t *testing.T) {
	checkErr(nil)
}
