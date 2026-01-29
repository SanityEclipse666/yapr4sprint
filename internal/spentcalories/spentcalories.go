package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	// Делим строку на слайс строк
	vals := strings.Split(data, ",")
	if len(vals) != 3 {
		return 0, "", 0, fmt.Errorf("неверный формат строки: требуется 3 элемента, получено %d", len(vals))
	}
	// Получаем количество шагов в int
	stepsStr := vals[0]
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка преобразования количества шагов: %w", err)
	}
	// Получаем вид активности в string
	activity := strings.TrimSpace(vals[1])
	// Получаем продолжительность в Time.Duration
	durationStr := vals[2]
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка преобразования продолжительности: %w", err)
	}
	// Возвращаем количество шагов, вид активности, продолжительность и nil для ошибки
	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	// Рассчитываем длину шага
	stepLength := height * stepLengthCoefficient
	// Рассчитываем дистанцию в метрах
	distanceMeters := float64(steps) * stepLength
	// Переводим дистанцию в километры
	distanceKM := distanceMeters / mInKm

	return distanceKM // Возвращаем дистанцию в километрах
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	// Если продолжительность меньше или равна нулю, возвращаем 0
	if duration <= 0 {
		return 0
	}
	// Вычисляем дистанцию
	dist := distance(steps, height)
	// Вычисляем среднюю скорость (дистанцию делим на время в часах)
	return dist / (duration.Seconds() / 3600)
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	// Получаем значения из строки данных
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err // Возвращаем пустую строку и ошибку
	}
	// Определяем тип тренировки и выполняем расчеты
	switch strings.ToLower(activity) {
	case "walking": // для ходьбы
		speed := meanSpeed(steps, height, duration) // Рассчитываем среднюю скорость
		distKM := distance(steps, height)           // Рассчитываем дистанцию
		calories, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
		return fmt.Sprintf(`Тип тренировки: Ходьба
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f ккал.`, duration.Hours(), distKM, speed, calories), nil

	case "running": // для бега
		speed := meanSpeed(steps, height, duration) // Рассчитываем среднюю скорость
		distKM := distance(steps, height)           // Рассчитываем дистанцию
		calories, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
		return fmt.Sprintf(`Тип тренировки: Бег
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f ккал.`, duration.Hours(), distKM, speed, calories), nil

	default:
		return "", fmt.Errorf("неизвестный тип тренировки") //  Обработка неизвестного типа тренировки.
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	// Проверка параметров на корректность
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть положительным")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть положительным")
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть положительным")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть больше нуля")
	}
	// Рассчитываем среднюю скорость
	speed := meanSpeed(steps, height, duration)
	// Рассчитываем количество калорий
	durationInMinutes := duration.Minutes()
	calories := (weight * speed * durationInMinutes) / minInH
	return calories, nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	// Проверка параметров на корректность
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть положительным")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть положительным")
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть положительным")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть больше нуля")
	}
	// Рассчитываем среднюю скорость
	speed := meanSpeed(steps, height, duration)
	// Рассчитываем количество калорий
	durationInMinutes := duration.Minutes()
	calories := (weight * speed * durationInMinutes) / minInH
	// Умножаем на корректирующий коэффициент
	calories = calories * walkingCaloriesCoefficient

	return calories, nil
}
