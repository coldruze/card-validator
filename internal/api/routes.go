package api

import "github.com/gorilla/mux"

// SetupRoutes настраивает маршруты и возвращает готовый маршрутизатор.
func SetupRoutes() *mux.Router {
	router := mux.NewRouter() // Создаем новый маршрутизатор с помощью пакета Gorilla Mux

	// Определяем маршрут и связываем его с обработчиком ValidateCardHandler
	router.HandleFunc("/valid", ValidateCardHandler)

	return router // Возвращаем настроенный маршрутизатор
}
