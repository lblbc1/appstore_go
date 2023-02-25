package repository

/*
厦门大学计算机专业 | 前华为工程师
专注《零基础学编程系列》  http://lblbc.cn/blog
包含：Java | 安卓 | 前端 | Flutter | iOS | 小程序 | 鸿蒙
公众号：蓝不蓝编程
*/
import (
	"gorm.io/gorm"
	"lblbc.cn/appStore/entity"
)

type Repository interface {
	QueryByCategory(categoryId string) []entity.AppInfo
	QueryById(id string) entity.AppInfo
	Search(keyword string) []entity.AppInfo
	QueryCategory() []entity.Category
}

type repoImpl struct {
	db *gorm.DB
}

func CreateRepository(db *gorm.DB) Repository {
	return &repoImpl{db: db}
}

func (repo *repoImpl) QueryByCategory(categoryId string) []entity.AppInfo {
	var apps []entity.AppInfo
	repo.db.Joins("left join appstore_app_category on appstore_app.id = appstore_app_category.App_id").Where("Category_id = ?", categoryId).Find(&apps)
	return apps
}

func (repo *repoImpl) QueryById(id string) entity.AppInfo {
	var data entity.AppInfo
	repo.db.Find(&data, id)
	return data
}

func (repo *repoImpl) Search(keyword string) []entity.AppInfo {
	var apps []entity.AppInfo
	repo.db.Where("Name like ?", "%"+keyword+"%").Find(&apps)
	return apps
}

func (repo *repoImpl) QueryCategory() []entity.Category {
	var categories []entity.Category
	repo.db.Find(&categories)
	return categories
}
