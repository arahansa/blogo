// mongotest
package main

import (
	"blogo/go/domain"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type Person struct {
	Name  string
	Phone string
}

type Counter struct {
	Seq int
}
type Counter2 int

func main() {
	session, err := mgo.Dial("mongodb://arahansa:1234@localhost:27017/blogo")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("blogo")
	peopleTable := db.C("people")
	err = peopleTable.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = peopleTable.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Phone:", result.Phone)
	articleTable := db.C("article")
	colCounter := db.C("counter")
	//--------자동증가 숫자를 수정해줘야함
	for i := 0; i < 3; i++ {

		change := mgo.Change{
			Update:    bson.M{"$inc": bson.M{"seq": 1}},
			ReturnNew: true,
		}
		doc := Counter{}
		colCounter.Find(bson.M{"_id": "article"}).Apply(change, &doc)
		fmt.Println(doc)
		//증가된 게시글에서 받은 id값을 주키로 넣어주고서 글 삽입

		err = articleTable.Insert(&domain.Article{Id: doc.Seq, Usernick: "arahansa", Subject: "hello!?", DayWrite: time.Now()})
		if err != nil {
			log.Fatal(err)
		}
		getArticle := domain.Article{}
		err = articleTable.Find(bson.M{"_id": doc.Seq}).One(&getArticle)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(getArticle)
	}

	fmt.Println("==== 리스트 === ")
	var articleList []domain.Article
	err = articleTable.Find(nil).Sort("-_id").All(&articleList)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(articleList)

}
