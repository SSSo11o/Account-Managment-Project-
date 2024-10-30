package repository

import (
	"account_managment/coinkeeper/internal/errs"
	"errors"
	"gorm.io/gorm"
)

func translateError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errs.ErrRecordNotFound
	}

	return err
}