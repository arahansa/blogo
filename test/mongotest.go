// mongotest
package main



import (
        "fmt"
	"log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
		"strconv"
)

type Person struct {
        Name string
        Phone string
}
type Article struct{
	Id int `bson:"_id,omitempty"`
	Usernick string `bson:"usernick"`
	Subject string `bson:"subject"`
}
func (a Article) String() string {
	return "나의 아이디 :"+strconv.Itoa(a.Id)+", 닉네임 :"+a.Usernick+",글제목:"+a.Subject
}
type Counter struct{
	Seq int
}


func main() {
        session, err := mgo.Dial("mongodb://arahansa:1234@localhost:27017/blogo")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)
		db:=session.DB("blogo")
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
		
		//--------자동증가 숫자를 수정해줘야함
		for i:=0;i<3;i++{
		boardCounter:=db.C("counters")
		change := mgo.Change{
        	Update: bson.M{"$inc": bson.M{"seq": 1}},
        	ReturnNew: true,
		}
		doc := Counter{}
		boardCounter.Find(bson.M{"_id": "userid"}).Apply(change, &doc)
		fmt.Println(doc.Seq)
		//증가된 게시글에서 받은 id값을 주키로 넣어주고서 글 삽입
		articleTable := db.C("article")
		err = articleTable.Insert(&Article{doc.Seq, "arahansa", "hello!?"})
		if err != nil {
                log.Fatal(err)
        }
		getArticle := Article{}
		err = articleTable.Find(bson.M{"_id":doc.Seq}).One(&getArticle)
		if err != nil {
                log.Fatal(err)
        }
		fmt.Println(getArticle)
		}
		
}
