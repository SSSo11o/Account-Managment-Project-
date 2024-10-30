package controllers

import (
	"account_managment/coinkeeper/internal/errs"
	"account_managment/coinkeeper/internal/models"
	"account_managment/coinkeeper/internal/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllCards(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	cards, err := service.GetAllCards(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cards": cards})
}

func GetCardByID(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	cardID, err := strconv.Atoi(c.Param("cardID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid card ID"})
		return
	}
	card, err := service.GetCardByID(userID, uint(cardID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, card)
}

func CreateCard(c *gin.Context) {
	var card models.Card

	if err := c.BindJSON(&card); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	card.UserID = userID // Устанавливаем ID пользователя

	if err := service.CreateCard(card); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, defaultResponse{Message: "Card created successfully"})
}


func UpdateCardBalance(c *gin.Context) {
	var updateRequest struct {
		CardID uint    `json:"card_id"`
		Amount float32 `json:"amount"` // Сумма для пополнения
	}

	if err := c.BindJSON(&updateRequest); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.UpdateCardBalance(updateRequest.CardID, updateRequest.Amount); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, defaultResponse{Message: "Card balance updated successfully"})
}


func DeleteCard(c *gin.Context) {
	cardID, err := strconv.Atoi(c.Param("id")) // Получаем ID карты из параметров URL
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	userID := c.GetUint(userIDCtx) // Получаем userID из контекста
	if userID == 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.DeleteCard(uint(cardID), userID); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, defaultResponse{Message: "Card deleted successfully"})
}
