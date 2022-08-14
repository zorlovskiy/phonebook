package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/gorilla/mux"
	"github.com/zorlovskiy/phonebook/config"
	"github.com/zorlovskiy/phonebook/domain"

	postgresRepo "github.com/zorlovskiy/phonebook/repository/postgres"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// чтение конфига псмотри либу caarlos/env
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg)
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

	db.AutoMigrate(&domain.Contact{})

	// инит структуры которая работает с базой
	contactStore := postgresRepo.NewContactStore(db)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		var contact domain.Contact

		err := json.NewDecoder(r.Body).Decode(&contact)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = contactStore.Create(&contact)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}).Methods("POST")
	// обновить данные контакта
	router.HandleFunc("/contact/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		var updateContact domain.Contact

		err := json.NewDecoder(r.Body).Decode(&updateContact)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err = contactStore.Update(&updateContact); err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}

	}).Methods("PUT")

	// поиск по имени
	router.HandleFunc("/contact/{fname}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		nameParams := vars["fname"]

		type GetByFNameResponse struct {
			Contacts []domain.Contact `json:"contacts"`
		}

		var response GetByFNameResponse

		response.Contacts, err = contactStore.GetByFName(nameParams)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}).Methods("GET")

	// поиск по номеру
	router.HandleFunc("/contact/bynumber/{number}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		numbParams := vars["number"]

		type GetByNumberResponse struct {
			Contacts []domain.Contact `json:"contacts"`
		}

		var response GetByNumberResponse

		response.Contacts, err = contactStore.GetByNumber(numbParams)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}).Methods("GET")

	//удаление по id
	router.HandleFunc("/contact/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idDelete := vars["id"]

		err := contactStore.DeleteByID(idDelete)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")

	}).Methods("DELETE")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
