// ArticleRepository
package repository

import (
	"blogo/go/domain"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

var (
	session                mgo.Session
	db                     *mgo.Database
	colArticle, colCounter *mgo.Collection
)

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

func GetOneArticle(id int) domain.Article {
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

func DeleteArticle(id int) {
	colArticle.RemoveId(id)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
