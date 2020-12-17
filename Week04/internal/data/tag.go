package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"time"
)

type Tag struct {
	Model

	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`

}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func GetTags(pageNum int, pageSize int, maps interface {}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface {}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagName(name string)  bool {
	var tag Tag
	db.Select("id").Where("name = ?",name).First(&tag)
	return tag.ID > 0
}
func ExistTagId(id int)  bool {
	var tag Tag
	db.Select("id").Where("id = ?",id).First(&tag)
	return tag.ID > 0
}

func AddTag(name string, state int, createdBy string) bool{
	db.Create(&Tag {
		Name : name,
		State : state,
		CreatedBy : createdBy,
	})

	return true
}

func UpdateTag(id int, data interface{}) error{
	err := db.Model(&Tag{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return errors.Wrap(err,"更新文章失败")
	}
	return nil
}

func DeleteTag(id int) bool{
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

func CleanAllTag() bool  {
	db.Unscoped().Where("deleted_on != ?", 0).Delete(&Tag{})
	return true
}
