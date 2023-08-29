package api

import (
	"encoding/json"
	"fmt"
	"github.com/coldruze/card-validator/pkg/valid"
	"log"
	"net/http"
)

// PersonData представляет структуру данных для информации о человеке.
type PersonData struct {
	Name       string
	CardNumber string
}

// ValidateCardHandler обрабатывает запросы для проверки номера карты.
func ValidateCardHandler(w http.ResponseWriter, r *http.Request) {
	var (
		p PersonData
	)
	const (
		BadRequest          = http.StatusBadRequest
		InternalServerError = http.StatusInternalServerError
	)

	const InvalidCardMsg = "Неверный номер карты"

	// Декодируем JSON-данные из тела запроса в структуру PersonData
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		handleError(w, err, BadRequest)
		return
	}

	// Проверяем валидность номера карты с помощью функции IsValidCardNum
	if valid.IsValidCardNum(p.CardNumber) {
		var paymentSystem string

		//Определяем платёжную систему
		switch p.CardNumber[0:1] {
		case "2":
			paymentSystem = "Мир"
		case "3":
			paymentSystem = "American Express"
		case "4":
			paymentSystem = "VISA"
		case "5":
			paymentSystem = "MasterCard"
		}

		// Если номер карты валиден, возвращаем информацию о платёжной системе
		fmt.Fprintln(w, "Платежная система:", paymentSystem)

		//Возвращаем информацию о человеке в формате JSON
		err := json.NewEncoder(w).Encode(p)
		if err != nil {
			handleError(w, err, InternalServerError)
			return
		}
	} else {
		// Если номер карты невалиден, возвращаем сообщение об ошибке
		fmt.Fprint(w, InvalidCardMsg)
	}
}

// handleError обрабатывает ошибки, отправляя соответствующий HTTP-ответ с ошибкой.
func handleError(w http.ResponseWriter, err error, statusCode int) {
	log.Printf("Error: %v", err)
	http.Error(w, err.Error(), statusCode)
}
