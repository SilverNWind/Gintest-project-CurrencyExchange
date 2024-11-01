package controllers

import (
	"errors"
	"exchangeapp/global"
	"exchangeapp/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateExchangeRates(ctx *gin.Context) {
	var exchangeRates models.ExchangeRates

	if err := ctx.ShouldBindJSON(&exchangeRates); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	exchangeRates.Date = time.Now()

	if err := global.Db.AutoMigrate(&exchangeRates); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "description": "Failed to migarate."})
	}

	if err := global.Db.Create(&exchangeRates).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "description": "Failed to create the table."})
	}

	ctx.JSON(http.StatusOK, exchangeRates)

}

func GetExchangeRates(ctx *gin.Context) {
	var exchangeRates []models.ExchangeRates

	if err := global.Db.Find(&exchangeRates).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		return
	}
	ctx.JSON(http.StatusOK, exchangeRates)

}
