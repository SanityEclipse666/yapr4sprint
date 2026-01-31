package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	// Делим строку на слайс строк
	vals := strings.Split(data, ",")
	if len(vals) != 2 {
		return 0, 0, fmt.Errorf("invalid input string format: expected two elements, have %d", len(vals))
	}
	// Преобразуем первый элемент в слайсе (количество шагов) в int
	stepsStr := vals[0]
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, 0, fmt.Errorf("steps count conversion error: %w", err)
	}
	// Проверям: количество шагов должно быть больше нуля
	if steps <= 0 {
		return 0, 0, errors.New("steps count must be more than 0")
	}
	// Преобразуем второй элемент слайса в time.Duration
	durationStr := vals[1]
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, 0, fmt.Errorf("duration conversion error: %w", err)
	}
	// Проверяем: длительность должна быть больше нуля
	if duration <= 0 {
		return 0, 0, errors.New("duration must be more than 0")
	}
	// Далее если все в порядке - возвращаем значения
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	// Получаем данные о количестве шагов и продолжительности прогулки
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		log.Println()
		return ""
	}
	// Проверка положительного количества шагов
	if steps <= 0 {
		return ""
	}
	// Вычисляем дистанцию в метрах
	distanceMeters := float64(steps) * stepLength
	// Переводим дистанцию в километры
	distanceKM := distanceMeters / mInKm
	// Вычисляем колистве калорий, потраченных на прогулке
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Println("calculating calories error:", err)
		return ""
	}
	// Создаем строку ответа
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distanceKM, calories)
}
