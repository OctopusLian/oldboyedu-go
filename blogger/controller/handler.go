package controller

import (
	"fmt"
	"log"
	"net/http"
	"oldboyedu-go/blogger/service"
	"oldboyedu-go/blogger/utils"
	"path"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func IndexHandle(c *gin.Context) {
	articleRecordList, err := service.GetArticleRecordList(0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	allCategoryList, err := service.GetAllCategoryList()
	if err != nil {
		fmt.Printf("get category list failed, err:%v\n", err)
	}
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": allCategoryList,
	})
}

func CategoryList(c *gin.Context) {
	categoryIdStr := c.Query("category_id")
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	articleRecordList, err := service.GetArticleRecordListById(int(categoryId), 0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	allCategoryList, err := service.GetAllCategoryList()
	if err != nil {
		fmt.Printf("get category list failed, err:%v\n", err)
	}
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": allCategoryList,
	})
}

func NewArticle(c *gin.Context) {
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		fmt.Printf("get article failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/post_article.html", categoryList)
}

func LeaveNew(c *gin.Context) {
	leaveList, err := service.GetLeaveList()
	if err != nil {
		fmt.Printf("get leave failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/gbook.html", leaveList)
}

func AboutMe(c *gin.Context) {
	c.HTML(http.StatusOK, "views/about.html", gin.H{
		"title": "Posts",
	})
}

func ArticleSubmit(c *gin.Context) {
	content := c.PostForm("content")
	author := c.PostForm("author")
	categoryIdStr := c.PostForm("category_id")
	title := c.PostForm("title")
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	err = service.InsertArticle(content, author, title, categoryId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}

func ArticleDetail(c *gin.Context) {
	articleIdStr := c.Query("article_id")
	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	articleDetail, err := service.GetArticleDetail(articleId)
	if err != nil {
		fmt.Printf("get article detail failed,article_id:%d err:%v\n", articleId, err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	relativeArticle, err := service.GetRelativeAricleList(articleId)
	if err != nil {
		fmt.Printf("get relative article failed, err:%v\n", err)
	}
	prevArticle, nextArticle, err := service.GetPrevAndNextArticleInfo(articleId)
	if err != nil {
		fmt.Printf("get prev or next article failed, err:%v\n", err)
	}
	allCategoryList, err := service.GetAllCategoryList()
	if err != nil {
		fmt.Printf("get all category failed, err:%v\n", err)
	}
	commentList, err := service.GetCommentList(articleId)
	if err != nil {
		fmt.Printf("get comment list failed, err:%v\n", err)
	}
	var m map[string]interface{} = make(map[string]interface{}, 10)
	m["detail"] = articleDetail
	m["relative_article"] = relativeArticle
	m["prev"] = prevArticle
	m["next"] = nextArticle
	m["category"] = allCategoryList
	m["article_id"] = articleId
	m["comment_list"] = commentList
	c.HTML(http.StatusOK, "views/detail.html", m)
}

func UploadFile(c *gin.Context) {
	// single file
	file, err := c.FormFile("upload")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	log.Println(file.Filename)
	rootPath := utils.GetRootDir()
	u2 := uuid.NewV4()
	if err != nil {
		return
	}
	ext := path.Ext(file.Filename)
	url := fmt.Sprintf("/static/upload/%s%s", u2, ext)
	dst := filepath.Join(rootPath, url)
	// Upload the file to specific dst.
	_ = c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"uploaded": true,
		"url":      url,
	})
}

func CommentSubmit(c *gin.Context) {
	comment := c.PostForm("comment")
	author := c.PostForm("author")
	email := c.PostForm("email")
	articleIdStr := c.PostForm("article_id")
	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	err = service.InsertComment(comment, author, email, articleId)
	if err != nil {
		fmt.Printf("insert comment failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	url := fmt.Sprintf("/article/detail/?article_id=%d", articleId)
	c.Redirect(http.StatusMovedPermanently, url)
}

func LeaveSubmit(c *gin.Context) {
	comment := c.PostForm("comment")
	author := c.PostForm("author")
	email := c.PostForm("email")
	err := service.InsertLeave(author, email, comment)
	if err != nil {
		fmt.Printf("insert leave failed, err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	url := fmt.Sprintf("/leave/new/")
	c.Redirect(http.StatusMovedPermanently, url)
}
