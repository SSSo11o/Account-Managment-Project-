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

func GetAllIncome(c *gin.Context) {
	query := c.Query("q")

	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	income, err := service.GetAllIncome(userID, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"income": income})
}

func GetIncomeByID(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	incomeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	income, err := service.GetIncomeByID(userID, uint(incomeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, income)
}

func CreateIncome(c *gin.Context) {
	var income models.Income
	if err := c.BindJSON(&income); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	income.UserID = userID
	if err := service.CreateIncome(income); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, defaultResponse{Message: "income created successfully"})
}

func UpdateIncome(c *gin.Context) {
	incomeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	var income models.Income
	if err = c.BindJSON(&income); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	income.ID = uint(incomeID)
	income.UserID = userID
	if err = service.UpdateIncome(income); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, defaultResponse{Message: "income updated successfully"})
}

func DeleteIncome(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		handleError(c, errs.ErrValidationFailed)
		return
	}
	incomeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if err = service.DeleteIncome(incomeID, uint(userID)); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, defaultResponse{Message: "income deleted successfully"})
}

func GetTotalIncome(c *gin.Context) {
	userID := c.GetUint(userIDCtx)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	totalIncome, err := service.GetTotalIncome(userID)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_income": totalIncome})
}

func GetFilteredIncome(c *gin.Context) {
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

	income, err := service.GetFilteredIncome(userID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"income": income})
}
