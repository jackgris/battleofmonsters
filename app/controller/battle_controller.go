package controller

import (
	"battle-of-monsters/app/db"
	"battle-of-monsters/app/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListBattles(context *gin.Context) {
	var battle []models.Battle

	var result *gorm.DB

	if result = db.CONN.Find(&battle); result.Error != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	log.Printf("Found %v battles", result.RowsAffected)
	context.JSON(http.StatusOK, &battle)
}

func DeleteBattle(context *gin.Context) {
	battleID := context.Param("battleID")

	var battle models.Battle

	if result := db.CONN.First(&battle, battleID); result.Error != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	if result := db.CONN.Delete(&models.Monster{}, battleID); result.Error != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	context.Status(http.StatusNoContent)
}
