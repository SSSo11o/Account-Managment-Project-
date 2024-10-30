package service

import (
	"account_managment/coinkeeper/internal/errs"
	"account_managment/coinkeeper/internal/models"
	"account_managment/coinkeeper/internal/pkg/repository"
	"errors"
	"time"
)

func GetAllOutcome(userID uint, query string) (outcome []models.Outcome, err error) {
	outcome, err = repository.GetAllOutcome(userID, query)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func GetOutcomeByID(userID, outcomeID uint) (outcome models.Outcome, err error) {
	outcome, err = repository.GetOutcomeByID(userID, outcomeID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return outcome, errs.ErrOperationNotFound
		}
		return outcome, err
	}
	return outcome, nil
}

func CreateOutcome(outcome models.Outcome) error {
	if err := repository.CreateOutcome(outcome); err != nil {
		return err
	}
	return nil
}

func UpdateOutcome(outcome models.Outcome) error {
	if err := repository.UpdateOutcome(outcome); err != nil {
		return err
	}
	return nil
}

func DeleteOutcome(outcomeID int, userID uint) error {
	if err := repository.DeleteOutcome(uint(outcomeID), userID); err != nil {
		return err
	}
	return nil
}
func GetTotalOutcome(userID uint) (float64, error) {
	totalIncome, err := repository.GetTotalOutcome(userID)
	if err != nil {
		return 0, err
	}
	return totalIncome, nil
}

func GetIncomeAndOutcomeReport(userID uint) (float64, float64, error) {
	totalIncome, err := GetTotalIncome(userID)
	if err != nil {
		return 0, 0, err
	}

	totalOutcome, err := GetTotalOutcome(userID)
	if err != nil {
		return 0, 0, err
	}

	return totalIncome, totalOutcome, nil
}

func GetIncomeAndOutcomeBalance(userID uint) (float64, float64, float64, error) {
	totalIncome, err := GetTotalIncome(userID)
	if err != nil {
		return 0, 0, 0, err
	}

	totalOutcome, err := GetTotalOutcome(userID)
	if err != nil {
		return 0, 0, 0, err
	}

	balance := totalIncome - totalOutcome
	return totalIncome, totalOutcome, balance, nil
}

func GetFilteredOutcome(userID uint, startDate, endDate time.Time) ([]models.Outcome, error) {
	return repository.GetOutcomeByFilter(userID, startDate, endDate)
}
