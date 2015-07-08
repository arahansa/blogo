// ArticleRepository
package repository

import (
	"blogo/go/domain"
	"fmt"
	"log"
	"strconv"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	session                mgo.Session
	db                     *mgo.Database
	colArticle, colCounter *mgo.Collection
)
const COUNT_PER_PAGE = 10;

type Counter struct {
	Seq int
}

func init() {
	fmt.Println("DB start!!")
	session, err := mgo.Dial("mongodb://arahansa:1234@localhost:27017/blogo")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	db = session.DB("blogo")
	colArticle = db.C("article")
	colCounter = db.C("counter")
	fmt.Println("DB init!!")
	defer fmt.Println("Close?!")
}

func GetArticleList() []domain.Article {
	var articles []domain.Article
	err := colArticle.Find(nil).Sort("-_id").All(&articles)
	checkErr(err)
	return articles
}

func GetArticleListForPage(requestPage int) ([]domain.Article, domain.PageInfo){
	if requestPage==0 {requestPage++}
	totalArticleCount := Count()
	//Paging 알고리즘 :: 도메인 영역으로 빼야할듯 싶다.
	totalPageCount := totalArticleCount / COUNT_PER_PAGE;
    if (totalArticleCount % COUNT_PER_PAGE) != 0{
    	totalPageCount++;
    }
    beginPage  := (requestPage - 1) / COUNT_PER_PAGE * COUNT_PER_PAGE + 1
    endPage := beginPage + (COUNT_PER_PAGE-1)
	if endPage > totalPageCount{
		endPage = totalPageCount
	}
		
	var pageinfo  domain.PageInfo
	pageinfo.BeginPage = beginPage
	pageinfo.EndPage = endPage
	pageinfo.TotalPageCount = totalPageCount
	log.Println("카운트 :", totalArticleCount, ", Pageinfo:", pageinfo)
	//mongo 디비에서는 그냥 skip 으로 넘길 내용을 계산함^^;
	skip := COUNT_PER_PAGE * (requestPage -1)
	var articles []domain.Article
	colArticle.Find(nil).Sort("-_id").Limit(COUNT_PER_PAGE).Skip(skip).All(&articles)
	
	return articles, pageinfo
}

func GetOneArticle(id int) domain.Article {
	//조회수 증가
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"numread": 1}},
		ReturnNew: true,
	}
	colArticle.Find(bson.M{"_id": id}).Apply(change, nil)
	//글 읽어오기
	result := domain.Article{}
	colArticle.Find(bson.M{"_id": id}).One(&result)
	return result
}
func Save(article *domain.Article) (int, error) {
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"seq": 1}},
		ReturnNew: true,
	}
	counter := Counter{}
	colCounter.Find(bson.M{"_id": "article"}).Apply(change, &counter)
	article.Id = counter.Seq
	article.DayWrite = time.Now()
	colArticle.Insert(article)
	return article.Id, nil
}

func Update(article *domain.Article) (int, error) {
	colArticle.Update(bson.M{"_id": article.Id}, bson.M{"$set": bson.M{"subject": article.Subject, "content": article.Content}})
	return article.Id, nil
}

func Count() int {
	count, err := colArticle.Count()
	checkErr(err)
	return count
}

func DeleteArticle(id int) {
	colArticle.RemoveId(id)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CreateDummyData() {
	var article domain.Article
	for i := 0; i < 30; i++ {
		article = domain.Article{}
		article.Usernick = "아라한사"
		article.Subject = "안녕하세요" + strconv.Itoa(i)
		article.Content = "헬로월드~~" + strconv.Itoa(i)
		Save(&article)
	}
}

func RemoveAll(){
	colArticle.RemoveAll(nil)
}
