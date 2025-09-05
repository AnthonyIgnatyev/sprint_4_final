package spentcalories

import (
	"fmt"
	"strconv"
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
	dataStororage := make([]string, 3)
	for _, value := range data {
		dataStororage = append(dataStororage, string(value))
	}

	if len(dataStororage) != 3 {
		return 0, "", 0, fmt.Errorf("Неверное количество данных для обработки")
	}

	steps, err := strconv.Atoi(dataStororage[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("Ошибка преобразования строки в число шагов")
	}

	duration, err := time.ParseDuration(dataStororage[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("Ошибка преобразования строки в длительность")
	}

	return steps, dataStororage[1], duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := stepLengthCoefficient * height
	distance := float64(steps) * stepLength / mInKm
	return distance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration == 0 {
		return 0
	} else {
		return float64(steps) / (duration.Hours() / minInH)
	}
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	var message string
	steps, trainingType, duration, err := parseTraining(data)
	if err != nil {
		return "", err
	}
	switch trainingType {
	case "Бег":
		message += fmt.Sprintf("Тип тренировки: %s/n", trainingType)
		message += fmt.Sprintf("Длительность: %.2f ч./n", duration)
		message += fmt.Sprintf("Дистанция: %.2f км./n", distance(steps, height))
		message += fmt.Sprintf("Скоростьь: %.2f км/ч./n", meanSpeed(steps, height, duration))
		calories, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		message += fmt.Sprintf("Сожгли калорий: %.2f ккал./n", calories)
	case "Ходьба":
		message += fmt.Sprintf("Тип тренировки: %s/n", trainingType)
		message += fmt.Sprintf("Длительность: %.2f ч./n", duration)
		message += fmt.Sprintf("Дистанция: %.2f км./n", distance(steps, height))
		message += fmt.Sprintf("Скоростьь: %.2f км/ч./n", meanSpeed(steps, height, duration))
		calories, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		message += fmt.Sprintf("Сожгли калорий: %.2f ккал./n", calories)
	}
	return message, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps == 0 || duration == 0 {
		return 0, fmt.Errorf("Неверные данные")
	}
	speed := meanSpeed(steps, height, duration)
	durationInMin := duration.Minutes()
	calories := speed * durationInMin * weight / minInH

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps == 0 || duration == 0 {
		return 0, fmt.Errorf("Неверные данные")
	}
	speed := meanSpeed(steps, height, duration)
	durationInMin := duration.Minutes()
	calories := speed * durationInMin * weight / minInH
	calories *= walkingCaloriesCoefficient
	return calories, nil
}
