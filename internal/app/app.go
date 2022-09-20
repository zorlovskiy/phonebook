package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/zorlovskiy/phonebook/config"
	v1 "github.com/zorlovskiy/phonebook/internal/conroller/http/v1"
	"github.com/zorlovskiy/phonebook/internal/models"
	"github.com/zorlovskiy/phonebook/internal/usecase"
	"github.com/zorlovskiy/phonebook/internal/usecase/repository/pg"

	//httpserver "github.com/zorlovskiy/phonebook/pkg/httpserver/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Run(cfg *config.Config) {

	// соединение с БД
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, //значения при соединении берутся из полей структуры
		cfg.DBUser, //структура получает значения из переменных окружения, которые укажем при запуске программы
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
		cfg.SSLmode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		//fmt.Printf("%+v\n", err)
		log.Fatal(err)
	}

	phonebookUC := usecase.NewContactUseCase(
		pg.NewContactRepository(db),
	)

	err = db.AutoMigrate(&models.Contact{})
	if err != nil {
		//fmt.Printf("%+v\n", err)
		log.Fatal(err)
	}

	srv := &http.Server{
		Handler:      v1.Router(phonebookUC),
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("http Server started")
	log.Fatal(srv.ListenAndServe())

}
