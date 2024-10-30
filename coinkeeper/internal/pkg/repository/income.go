package repository

import (
	"account_managment/coinkeeper/internal/db"
	"account_managment/coinkeeper/internal/logger"
	"account_managment/coinkeeper/internal/models"
	"time"
)

func GetAllIncome(userID uint, query string) ([]models.Income, error) {
	var income []models.Income

	query = "%" + query + "%"

	err := db.GetDBConn().Model(&models.Income{}).
		Joins("JOIN users ON users.id = incomes.user_id").
		Where("incomes.user_id = ? AND incomes.is_deleted = false AND description iLIKE ?", userID, query).
		Order("incomes.id").
		Find(&income).Error
	if err != nil {
		logger.Error.Println("[repository.GetAllIncome] cannot get all income. Error is:", err.Error())
		return nil, translateError(err)
	}
	return income, nil
}

func GetIncomeByID(userID, incomeID uint) (income models.Income, err error) {
	err = db.GetDBConn().Model(&models.Income{}).
		Joins("JOIN users ON users.id = incomes.user_id").
		Where("incomes.user_id = ? AND is_deleted = false AND incomes.id = ?", userID, incomeID).
		First(&income).Error
	if err != nil {
		logger.Error.Println("[repository.GetIncomeByID] cannot get income by id. Error is:", err.Error())
		return models.Income{}, translateError(err)
	}
	return income, nil
}

func CreateIncome(income models.Income) error {
	err := db.GetDBConn().Create(&income).Error
	if err != nil {
		logger.Error.Println("[repository.CreateIncome] cannot create income. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}

func UpdateIncome(income models.Income) error {
	err := db.GetDBConn().Model(&income).Where("id = ?", income.ID).Updates(income).Error

	if err != nil {
		logger.Error.Println("[repository.UpdateIncome] cannot update income. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}

func DeleteIncome(incomeID int, userID uint) error {
	err := db.GetDBConn().
		Model(&models.Income{}).
		Where("id = ? AND user_id = ?", incomeID, userID).
		Update("is_deleted", true).Error
	if err != nil {
		logger.Error.Println("[repository.DeleteIncome] cannot delete income. Error is:", err.Error())
		return err
	}
	return nil

}

func GetTotalIncome(userID uint) (float64, error) {
	var totalIncome float64
	err := db.GetDBConn().
		Model(&models.Income{}).
		Where("user_id = ? AND is_deleted = false", userID).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&totalIncome).Error
	if err != nil {
		logger.Error.Println("[repository.GetTotalIncome] cannot get total income. Error is:", err.Error())
		return 0, translateError(err)
	}
	return totalIncome, nil
}

func GetIncomeByFilter(userID uint, startDate, endDate time.Time) ([]models.Income, error) {
	var income []models.Income
	dbConn := db.GetDBConn().Model(&models.Income{}).Where("user_id = ? AND is_deleted = false", userID)
	if !startDate.IsZero() && !endDate.IsZero() {
		dbConn = dbConn.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}
	err := dbConn.Order("created_at").Find(&income).Error
	if err != nil {
		return nil, err
	}
	return income, nil
}
