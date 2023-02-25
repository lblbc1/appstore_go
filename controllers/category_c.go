package controllers

/*
厦门大学计算机专业 | 前华为工程师
专注《零基础学编程系列》  http://lblbc.cn/blog
包含：Java | 安卓 | 前端 | Flutter | iOS | 小程序 | 鸿蒙
公众号：蓝不蓝编程
*/
import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lblbc.cn/appStore/entity"
	"lblbc.cn/appStore/helper"
	"lblbc.cn/appStore/repository"
)

type CategoryController interface {
	QueryCategory(c *gin.Context)
}

type categoryController struct {
	repo repository.Repository
}

func NewCategoryController(repo repository.Repository) CategoryController {
	return &categoryController{repo: repo}
}

func (c *categoryController) QueryCategory(ctx *gin.Context) {
	var data []entity.Category = c.repo.QueryCategory()
	response := helper.SuccessResponse(0, "", data)
	ctx.JSON(http.StatusOK, response)
}
