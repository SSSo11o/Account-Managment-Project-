package controllers

import (
	"account_managment/coinkeeper/internal/errs"
	"account_managment/coinkeeper/internal/models"
	"account_managment/coinkeeper/internal/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetAllOutcome(c *gin.Context) {
	query := c.Query("q")

	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	outcome, err := service.GetAllOutcome(userID, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"outcome": outcome})
}

func GetOutcomeByID(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	outcomeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	outcome, err := service.GetOutcomeByID(userID, uint(outcomeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, outcome)
}

func CreateOutcome(c *gin.Context) {
	var outcome models.Outcome
	if err := c.BindJSON(&outcome); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	outcome.UserID = userID
	if err := service.CreateOutcome(outcome); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, defaultResponse{Message: "outcome created successfully"})
}

func UpdateOutcome(c *gin.Context) {
	outcomeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	var outcome models.Outcome
	if err = c.BindJSON(&outcome); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	outcome.ID = uint(outcomeID)

	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	outcome.ID = uint(outcomeID)
	outcome.UserID = userID
	if err = service.UpdateOutcome(outcome); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, defaultResponse{Message: "outcome updated successfully"})
}

func DeleteOutcome(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	outcomeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if err = service.DeleteOutcome(outcomeID, uint(userID)); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, defaultResponse{Message: "outcome deleted successfully"})
}

func GetTotalOutcome(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	totalOutcome, err := service.GetTotalOutcome(userID)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_outcome": totalOutcome})
}

func GetIncomeAndOutcomeBalanceReport(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	totalIncome, totalOutcome, balance, err := service.GetIncomeAndOutcomeBalance(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_income":  totalIncome,
		"total_outcome": totalOutcome,
		"balance":       balance,
	})
}

func GetFilteredOutcome(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	var startDate, endDate time.Time
	var err error

	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат start_date"})
			return
		}
	}

	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат end_date"})
			return
		}
	}

	outcome, err := service.GetFilteredOutcome(userID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"outcome": outcome})
}
