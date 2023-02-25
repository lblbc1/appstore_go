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

type AppController interface {
	QueryByCategory(c *gin.Context)
	QueryById(c *gin.Context)
	Search(c *gin.Context)
}

type appController struct {
	repo repository.Repository
}

func NewAppController(repo repository.Repository) AppController {
	return &appController{repo: repo}
}

func (c *appController) QueryById(ctx *gin.Context) {
	id := ctx.Param("id")
	var data entity.AppInfo = c.repo.QueryById(id)
	response := helper.SuccessResponse(0, "", data)
	ctx.JSON(http.StatusOK, response)
}

func (c *appController) Search(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	var data []entity.AppInfo = c.repo.Search(keyword)
	response := helper.SuccessResponse(0, "", data)
	ctx.JSON(http.StatusOK, response)
}

func (c *appController) QueryByCategory(ctx *gin.Context) {
	categoryId := ctx.Query("categoryId")
	var data []entity.AppInfo = c.repo.QueryByCategory(categoryId)
	response := helper.SuccessResponse(0, "", data)
	ctx.JSON(http.StatusOK, response)
}
