package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Model
	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// 查询是否存在指定Id的文章
func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

// 获取文章总数
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

// 获取符合条件的文章（分页）
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	/*
		此处的preload()方法是一个预加载器,他会执行两条SQL,分别是:
		SELECT * FROM blog_article;
		SELECT * FROM blog_tags WHERE id IN (1,2,3,4);
		那么在查询出结构后，gorm内部处理对应的映射逻辑，将其填充到Article的Tag中，
		会特别方便，并且避免了循环查询
	*/
	return
}

// 根据id获取指定的文章
func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	/*
		gorm 是怎样内部将Tag绑定到对应的Article上的:
		Article有一个结构体成员是TagID，就是外键。
		gorm会通过类名+ID 的方式去找到这两个类之间的关联关系
		Article有一个结构体成员是Tag，
		就是我们嵌套在Article里的Tag结构体，我们可以通过Related进行关联查询
	*/
	return
}

// 编辑文章
func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

// 添加一个文章
func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})

	return true
}

// 删除一篇文章
func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}

// 创建时间
func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

// 更新时间
func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
