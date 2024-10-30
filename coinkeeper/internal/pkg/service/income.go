package service

import (
	"account_managment/coinkeeper/internal/errs"
	"account_managment/coinkeeper/internal/models"
	"account_managment/coinkeeper/internal/pkg/repository"
	"errors"
	"time"
)

func GetAllIncome(userID uint, query string) (income []models.Income, err error) {
	income, err = repository.GetAllIncome(userID, query)
	if err != nil {
		return nil, err
	}
	return income, nil
}

func GetIncomeByID(userID, incomeID uint) (income models.Income, err error) {
	income, err = repository.GetIncomeByID(userID, incomeID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return income, errs.ErrOperationNotFound
		}
		return models.Income{}, err
	}
	return income, nil
}

func CreateIncome(income models.Income) error {
	if err := repository.CreateIncome(income); err != nil {
		return err
	}
	return nil
}

func UpdateIncome(income models.Income) error {
	if err := repository.UpdateIncome(income); err != nil {
		return err
	}
	return nil
}

func DeleteIncome(incomeID int, userID uint) error {
	if err := repository.DeleteIncome(incomeID, userID); err != nil {
		return err
	}
	return nil
}

func GetTotalIncome(userID uint) (float64, error) {
	totalIncome, err := repository.GetTotalIncome(userID)
	if err != nil {
		return 0, err
	}
	return totalIncome, nil
}

func GetFilteredIncome(userID uint, startDate, endDate time.Time) ([]models.Income, error) {
	return repository.GetIncomeByFilter(userID, startDate, endDate)
}
