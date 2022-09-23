package model

import (
	"ginblog/utils/errmsg"

	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// 新增文章
func CreateArt(data *Article) int {
	if err := db.Create(&data).Error; err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE

}

// TODO:查询分类下的所有文章
func GetCateArt(id, pageSize, pageNumber int) ([]Article, int, int) {
	var artList []Article
	var total int

	err := db.Preload("Category").Limit(pageSize).Offset((pageNumber-1)*pageSize).Where("cid = ?", id).Find(&artList).Count(&total).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return artList, errmsg.SUCCSE, total
}

// TODO:查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	if err := db.Preload("Category").Where("id = ?", id).First(&art).Error; err != nil {
		return art, errmsg.ERROR_ARTICLE_NOT_EXISIT
	}
	return art, errmsg.SUCCSE
}

// TODO:查询文章列表
func GetArt(pageSize, pageNumber int) ([]Article, int, int) {
	var artList []Article
	var total int
	err := db.Preload("Category").Limit(pageSize).Offset((pageNumber - 1) * pageSize).Find(&artList).Count(total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return artList, errmsg.SUCCSE, total
}

// 编辑分类信息
func EditArt(id int, data *Article) int {
	maps := map[string]interface{}{
		"title":   data.Title,
		"cid":     data.Cid,
		"desc":    data.Desc,
		"content": data.Content,
		"img":     data.Img,
	}
	if err := db.Model(&data).Where("id = ?", id).Updates(maps).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除分类
func DeleteArt(id int) int {
	var art Article
	if err := db.Where("id = ? ", id).Delete(&art).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
