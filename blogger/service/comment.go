package service

import (
	"oldboyedu-go/blogger/dao/db"
	"oldboyedu-go/blogger/model"
	"fmt"
	"time"
)

func InsertComment(comment, author, email string, articleId int64) (err error) {
	exist, err := db.IsArticleExist(articleId)
	if err != nil {
		fmt.Printf("query database failed, err:%v\n", err)
		return
	}
	if exist == false {
		err = fmt.Errorf("article id:%d not found", articleId)
		return
	}
	var c model.Comment
	c.ArticleId = articleId
	c.Content = comment
	c.Username = author
	c.CreateTime = time.Now()
	c.Status = 1
	err = db.InsertComment(&c)
	return
}

func GetCommentList(articleId int64) (commentList []*model.Comment, err error) {
	exist, err := db.IsArticleExist(articleId)
	if err != nil {
		fmt.Printf("query database failed, err:%v\n", err)
		return
	}
	if exist == false {
		err = fmt.Errorf("article id:%d not found", articleId)
		return
	}
	commentList, err = db.GetCommentList(articleId, 0, 100)
	return
}
