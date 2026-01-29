package daysteps

import (
	"fmt"
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
		return 0, 0, fmt.Errorf("неверный формат входной строки: ожидается два элемента, получено %d", len(vals))
	}
	// Преобразуем первый элемент в слайсе (количество шагов) в int
	stepsStr := vals[0]
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования количества шагов: %w", err)
	}
	// Проверям: количество шагов должно быть больше нуля
	if steps <= 0 {
		return 0, 0, fmt.Errorf("количество шагов должно быть больше 0")
	}
	// Преобразуем второй элемент слайса в time.Duration
	durationStr := vals[1]
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования продолжительности: %w", err)
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
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, distanceKM, duration)
	// Создаем строку ответа
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		steps, distanceKM, calories)
}
