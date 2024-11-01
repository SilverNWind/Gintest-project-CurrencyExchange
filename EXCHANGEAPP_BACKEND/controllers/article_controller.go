package controllers

import (
	"encoding/json"
	"errors"
	"exchangeapp/global"
	"exchangeapp/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var cacheKey = "articles"

func CreateArticles(ctx *gin.Context) {
	var article models.Article

	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := global.Db.AutoMigrate(&article); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "DB migration failed"})
	}

	if err := global.Db.Create(&article).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "DB creation failed"})
	}
	//缓存失效策略：主动更新缓存，待下次更新文章后未命中缓存会直接写入一份缓存
	if err := global.RedisDB.Del(cacheKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, article)

}

func GetArticles(ctx *gin.Context) {

	cachedData, err := global.RedisDB.Get(cacheKey).Result()

	//if the cached missed, fetch data from db and cache it
	if err == redis.Nil {
		var articles []models.Article

		if err := global.Db.Find(&articles).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
		articleJSON, err := json.Marshal(articles)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "JSON marshalling failed"})
			return
		}

		if err := global.RedisDB.Set(cacheKey, articleJSON, 10*time.Minute).Err(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, articles)

	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	} else {
		var articles []models.Article
		if err := json.Unmarshal([]byte(cachedData), &articles); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "JSON unmarshalling failed"})
			return
		}

		ctx.JSON(http.StatusOK, articles)
	}

}

func GetArticlesByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var article models.Article

	if err := global.Db.Where("id = ?", id).First(&article).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		return
	}

	ctx.JSON(http.StatusOK, article)
}
