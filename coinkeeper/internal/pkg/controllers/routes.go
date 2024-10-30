package controllers

import (
	"account_managment/coinkeeper/internal/configs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)

	r.GET("/ping", PingPong)

	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}

	apiG := r.Group("/api", checkUserAuthentication)

	incomeG := apiG.Group("/income")
	{
		incomeG.GET("", GetAllIncome)
		incomeG.POST("", CreateIncome)
		incomeG.GET("/:id", GetIncomeByID)
		incomeG.PUT("/:id", UpdateIncome)
		incomeG.DELETE("/:id", DeleteIncome)
		incomeG.GET("/filter", GetFilteredIncome)
	}

	outcomeG := apiG.Group("/outcome")
	{
		outcomeG.GET("", GetAllOutcome)
		outcomeG.POST("", CreateOutcome)
		outcomeG.GET("/:id", GetOutcomeByID)
		outcomeG.PUT("/:id", UpdateOutcome)
		outcomeG.DELETE("/:id", DeleteOutcome)
		outcomeG.GET("/filter", GetFilteredOutcome)
	}

	reportIncomeG := apiG.Group("/report_income")
	{
		reportIncomeG.GET("", GetTotalIncome)
	}

	reportOutcomeG := apiG.Group("/report_outcome")
	{
		reportOutcomeG.GET("", GetTotalOutcome)
	}

	reportIncomeAndOutcomeG := apiG.Group("/report_income_and_outcome")
	{
		reportIncomeAndOutcomeG.GET("", GetIncomeAndOutcomeBalanceReport)
	}

	cardG := apiG.Group("/cards")
	{
		cardG.GET("", GetAllCards)
		cardG.POST("", CreateCard)
		cardG.GET("/:id", GetCardByID)
		cardG.PUT("/:id", UpdateCardBalance)
		cardG.DELETE("/:id", DeleteCard)
	}

	userG := apiG.Group("/users")
	{
		userG.GET("", GetAllUsers)
	}
	return r
}

func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
