package db

import "account_managment/coinkeeper/internal/models"

func Migrate() error {
	err := dbConn.AutoMigrate(models.User{},
		models.Income{},
		models.Outcome{},
		models.Card{},
	)
	if err != nil {
		return err
	}
	return nil
}
