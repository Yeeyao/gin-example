package models

import "github.com/jinzhu/gorm"

// 增加 error 返回类型
// Article 文章
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

// ExistArticleByID check exist by ID
func ExistArticleByID(id int) (bool, error) {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true, nil
	}
	return false, nil
}

// GetArticleTotal 获取所有的 article
func GetArticleTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Article{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// GetArticles get article by page
func GetArticles(pageNum int, pageSize int, maps interface{}) ([]*Article, error) {
	var articles []*Article
	err := db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, nil
}

// GetArticle by id
func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	err = db.Model(&article).Related(&article.Tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &article, nil
}

// EditArticle 修改
func EditArticle(id int, data interface{}) error {
	if err := db.Model(&Article{}).Where("id = ?", id).Update(data).Error; err != nil {
		return err
	}
	return nil
}

// AddArticle 新增文章
func AddArticle(data map[string]interface{}) error {
	article := Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	}
	if err := db.Create(&article).Error; err != nil {
		return err
	}
	return nil
}

// DeleteArticle 删除文章
func DeleteArticle(id int) error {
	if err := db.Where("id = ?", id).Delete(Article{}).Error; err != nil {
		return err
	}
	return nil
}

// BeforeCreate 设置创建时间
// func (article *Article) BeforeCreate(scope *gorm.Scope) error {
// 	scope.SetColumn("CreatedOn", time.Now().Unix())
// 	return nil
// }

// BeforeUpdate 设置更新时间
// func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
// 	scope.SetColumn("modifiedOn", time.Now().Unix())
// 	return nil
// }

// CleanAllArticle 硬删除
func CleanAllArticle() error {
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{}).Error; err != nil {
		return err
	}
	return nil
}
