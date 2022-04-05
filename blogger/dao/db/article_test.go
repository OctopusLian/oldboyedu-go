/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-04-05 15:30:46
 * @LastEditors: neozhang
 * @LastEditTime: 2022-04-05 15:36:50
 */
package db

import (
	"oldboyedu-go/blogger/model"
	"testing"
	"time"
)

func init() {
	dns := "root:admin@tcp(localhost:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestInsertArticle(t *testing.T) {
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryId = 1
	article.ArticleInfo.CommentCount = 0
	article.Content = "this a test ak dkdkdkddkddkd"
	article.ArticleInfo.CreateTime = time.Now()
	article.ArticleInfo.Summary = `我是
									很多的
									内容`
	article.ArticleInfo.Title = "我是标题"
	article.ArticleInfo.Username = "Mr.Sun"
	article.ArticleInfo.ViewCount = 1
	article.Category.CategoryId = 1
	articleId, err := InsertArticle(article)
	if err != nil {
		t.Errorf("insert article failed, err:%v\n", err)
		return
	}
	t.Logf("insert article succ, articleId:%d\n", articleId)
}

func TestGetArticleList(t *testing.T) {
	articleList, err := GetArticleList(1, 15)
	if err != nil {
		t.Errorf("get article failed, err:%v\n", err)
		return
	}
	t.Logf("get article succ, len:%d\n", len(articleList))
}

func TestGetArticleInfo(t *testing.T) {
	articleInfo, err := GetArticleDetail(5)
	if err != nil {
		t.Errorf("get article failed, err:%v\n", err)
		return
	}
	t.Logf("get article succ, article:%#v\n", articleInfo)
}

func TestGetRelativeArticle(t *testing.T) {
	articleList, err := GetRelativeArticle(7)
	if err != nil {
		t.Errorf("get relative article failed, err:%v\n", err)
		return
	}
	for _, v := range articleList {
		t.Logf("id:%d title:%s\n", v.ArticleId, v.Title)
	}
}

func TestGetPrevArticleById(t *testing.T) {
	articelInfo, err := GetPrevArticleById(6)
	if err != nil {
		t.Errorf("get prev article failed, err:%v\n", err)
		return
	}
	t.Logf("artice info:%#v", articelInfo)
}

func TestGetNextArticleById(t *testing.T) {
	articelInfo, err := GetNextArticleById(6)
	if err != nil {
		t.Errorf("get prev article failed, err:%v\n", err)
		return
	}
	t.Logf("artice info:%#v", articelInfo)
}
