// Article
package domain

import (
	"strconv"
	"time"
)

type Article struct {
	Id       int       `bson:"_id,omitempty"`
	Usernick string    `bson:"usernick"`
	Subject  string    `bson:"subject"`
	Content  string    `bson:"content,omitempty`
	DayWrite time.Time `bson:"daywrite`
	NumRead  int
}

const (
	DATE_FORMAT     = "Jan _2, 2006"
	SQL_DATE_FORMAT = "2006-01-02"
)

func (a Article) String() string {
	return "\n나의 아이디 :" + strconv.Itoa(a.Id) + ", 닉네임 :" + a.Usernick +
		",글제목 : " + a.Subject + ",글쓴날짜 : " + a.DayWrite.Format(SQL_DATE_FORMAT) + ", 조회수 : " + strconv.Itoa(a.NumRead) +
		", 글내용 : " + a.Content
}
