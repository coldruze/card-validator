package main

import (
	"github.com/coldruze/card-validator/internal/api"
	"log"
	"net/http"
)

func main() {
	// Настройка маршрутов с помощью функции SetupRoutes из пакета "api"
	router := api.SetupRoutes()

	// Привязываем маршрутизатор к корневому пути "/"
	http.Handle("/", router)

	// Запускаем HTTP-сервер на адресе localhost:8000
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
