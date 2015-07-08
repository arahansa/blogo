// ArtcleRepository_test.go
package repository

import (
	"blogo/go/domain"
	"fmt"
	"testing"
)
var id int

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

func TestGetOutArticle(t *testing.T){
	getArticle := GetOneArticle(1000)
	fmt.Println("없는 게시글?:",getArticle)
}

func TestDeleteArticle(t *testing.T) {
	DeleteArticle(id)
	article := GetOneArticle(id)
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

func TestCount(t *testing.T){
	fmt.Println("게시글 개수 : ",Count())
}

func TestGetArticleListPage(t *testing.T){
	articles, pageinfo := GetArticleListForPage(1)
	fmt.Println("1페이지 게시글들 :",articles)
	fmt.Println("1페이지 페이지 정보 :", pageinfo)
}

func TestDummyAndRemoveAll(t *testing.T){
	CreateDummyData()
	RemoveAll()
}
