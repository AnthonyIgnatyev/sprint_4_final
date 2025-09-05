package daysteps

import (
	"fmt"
	"strconv"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	dataStorage := make([]string, 2)
	for _, value := range data {
		dataStorage = append(dataStorage, string(value))
	}
	if len(dataStorage) != 2 {
		return 0, 0, fmt.Errorf("Неверное количество данных для обработки")
	}
	steps, err := strconv.Atoi(dataStorage[0])
	if err != nil {
		return 0, 0, fmt.Errorf("Ошибка преобразования строки в число шагов")
	}
	duration, err := time.ParseDuration(dataStorage[1])
	if err != nil {
		return 0, 0, fmt.Errorf("Ошибка преобразования строки в длительность")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	var message string
	steps, duration, err := parsePackage(data)
	if err != nil {
		return ""
	} else if steps == 0 || duration == 0 {
		if steps == 0 {
			return ""
		}
	} else {
		message += fmt.Sprintf("Количество шагов: %d", steps)
	}

	distance := float64(steps) * stepLength / mInKm
	message += fmt.Sprintf("Дистанция составила %.2f км", distance)

	calories := WalkingSpentCalories()
	message += fmt.Sprintf("Вы сожгли %.2f ккал", calories)

	return message
}
