package repository

import (
	"account_managment/coinkeeper/internal/db"
	"account_managment/coinkeeper/internal/logger"
	"account_managment/coinkeeper/internal/models"
	"errors"
	"gorm.io/gorm"
	"time"
)

func GetAllOutcome(userID uint, query string) ([]models.Outcome, error) {
	var outcome []models.Outcome

	query = "%" + query + "%"

	err := db.GetDBConn().Model(&models.Outcome{}).
		Joins("JOIN users ON users.id = outcomes.user_id").
		Joins("JOIN outcome_categories ON outcome_categories.id = outcomes.category_id").
		Where("outcomes.user_id = ? AND outcomes.is_deleted = false AND outcomes.description iLIKE ?", userID, query).
		Order("outcomes.id").
		Find(&outcome).Error

	if err != nil {
		logger.Error.Println("[repository.GetAllOutcome] cannot get all outcome. Error is:", err.Error())
		return nil, translateError(err)
	}
	return outcome, nil

}

func GetOutcomeByID(userID, outcomeID uint) (models.Outcome, error) {
	var outcome models.Outcome

	err := db.GetDBConn().Model(&models.Outcome{}).
		Joins("JOIN outcome_categories ON outcome_categories.id = outcomes.category_id").
		Where("outcomes.id = ? AND outcomes.user_id = ? AND outcomes.is_deleted = false", outcomeID, userID).
		First(&outcome).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {

	}
	if err != nil {
		logger.Error.Println("[repository.GetOutcomeByID] cannot get outcome by id. Error is:", err.Error())
		return models.Outcome{}, translateError(err)
	}

	return outcome, nil
}

func CreateOutcome(outcome models.Outcome) error {
	err := db.GetDBConn().Create(&outcome).Error
	if err != nil {
		logger.Error.Println("[repository.CreateOutcome] cannot create outcome. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}

func UpdateOutcome(outcome models.Outcome) error {
	err := db.GetDBConn().Model(&outcome).Where("id = ?", outcome.ID).Save(outcome).Error

	if err != nil {
		logger.Error.Println("[repository.UpdateOutcome] cannot update outcome. Error is:", err.Error())
		return translateError(err)
	}
	return nil
}

func DeleteOutcome(outcomeID, userID uint) error {
	// Обновляем флаг is_deleted на true

	err := db.GetDBConn().Exec("UPDATE outcomes set is_deleted = true WHERE id = $1 AND user_id = $2", outcomeID, userID).Error
	if err != nil {
		logger.Error.Println("[repository.DeleteOutcome] cannot delete outcome. Error is:", err.Error())
		return translateError(err)
	}

	return nil
}

func GetTotalOutcome(userID uint) (float64, error) {
	var totalOutcome float64
	err := db.GetDBConn().
		Model(&models.Outcome{}).
		Where("user_id = ? AND is_deleted = false", userID).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&totalOutcome).Error
	if err != nil {
		logger.Error.Println("[repository.GetTotalOutcome] cannot get total income. Error is:", err.Error())
		return 0, translateError(err)
	}
	return totalOutcome, nil
}

func GetOutcomeByFilter(userID uint, startDate, endDate time.Time) ([]models.Outcome, error) {
	var outcomes []models.Outcome
	dbConn := db.GetDBConn().Model(&models.Outcome{}).Where("user_id = ? AND is_deleted = false", userID)
	if !startDate.IsZero() && !endDate.IsZero() {
		dbConn = dbConn.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}
	err := dbConn.Order("created_at").Find(&outcomes).Error
	if err != nil {
		return nil, err
	}
	return outcomes, nil
}
