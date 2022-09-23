package model

import (
	"ginblog/utils/errmsg"

	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key:auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 查询分类是否存在
func CheckCategory(name string) int {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCSE
}

// 新增分类
func CreateCategory(data *Category) int {
	if err := db.Create(&data).Error; err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE
}

// TODO:查询分类下的所有文章

// 查询分类列表
func GetCategory(pageSize, pageNumber int) ([]Category, int) {
	var cate []Category
	var total int
	err := db.Limit(pageSize).Offset((pageNumber - 1) * pageSize).Find(&cate).Count(total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}

// 编辑分类信息
func EditCategory(id int, data *Category) int {
	maps := map[string]interface{}{
		"name": data.Name,
	}
	if err := db.Model(&data).Where("id = ?", id).Updates(maps).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除分类
func DeleteCategory(id int) int {
	var cate Category
	if err := db.Where("id = ? ", id).Delete(&cate).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
