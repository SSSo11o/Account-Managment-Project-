package main

import (
	"account_managment/coinkeeper/internal/configs"
	"account_managment/coinkeeper/internal/db"
	"account_managment/coinkeeper/internal/logger"
	"account_managment/coinkeeper/internal/pkg/controllers"
	"account_managment/coinkeeper/internal/server"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := godotenv.Load("./cmd/.env"); err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %s", err)
	}

	if err := configs.ReadSettings(); err != nil {
		log.Fatalf("Ошибка чтения настроек: %s", err)
	}

	if err := logger.Init(); err != nil {
		log.Fatalf("Ошибка инициализации логгера: %s", err)
	}

	var err error
	err = db.ConnectToDB()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %s", err)
	}

	if err = db.Migrate(); err != nil {
		log.Fatalf("Ошибка миграции базы данных: %s", err)
	}

	mainServer := new(server.Server)
	go func() {
		if err = mainServer.Run(configs.AppSettings.AppParams.PortRun, controllers.InitRoutes()); err != nil {
			log.Printf("Ошибка при запуске HTTP сервера: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Printf("\n Начало завершение программ\n")

	//Close DB
	if sqlDB, err := db.GetDBConn().DB(); err == nil {
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Ошибка при закрытии соединения c БД: %s", err)
		}
	} else {
		log.Fatalf("Ошибка при получении *sql.DB из GORM: %s", err)
	}
	fmt.Println("Соединение c БД успешно закрыто")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = mainServer.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка при завершении работы сервера: %s", err)
	}

	fmt.Println("HTTP-сервис успешно выключен")
	fmt.Println("Конец завершения программы")
}
