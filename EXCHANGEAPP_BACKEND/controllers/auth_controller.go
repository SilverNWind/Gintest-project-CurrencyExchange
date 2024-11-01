package controllers

import (
	"exchangeapp/global"
	"exchangeapp/models"
	"exchangeapp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	//创建user实例
	var user models.User
	//绑定json数据到user实例
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//加密密码
	hashedPwd, err := utils.HashPassword(user.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "description": "bcrypt error"})
		return
	}
	//将加密的密码赋值给user实例中的密码字段
	user.Password = hashedPwd

	//使用JWT设置Token
	token, err := utils.GenerateJWT(user.Username)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "description": "JWT error"})
		return
	}
	//将user实例保存到数据库
	if err := global.Db.AutoMigrate(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "description": "database migrate error"})
		return
	}
	//创建user实例
	if err := global.Db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "description": "database create error"})
		return
	}
	//返回Token
	ctx.JSON(http.StatusOK, gin.H{"token": token})
	// ctx.JSON(http.StatusOK, user)
}

func Login(ctx *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	if err := global.Db.Where("username = ?", input.Username).Find(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong credentials"})
		return
	}

	if !utils.CheckPassword(input.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Worng credentials"})
		return
	}

	token, err := utils.GenerateJWT(user.Username)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
