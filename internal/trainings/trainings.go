package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/s444v/go-fifth-sprint/internal/personaldata"
	"github.com/s444v/go-fifth-sprint/internal/spentenergy"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	// ваш код ниже
	splited := strings.Split(datastring, ",")
	if len(splited) != 3 {
		return errors.New("string format does not match")
	}
	steps, err := strconv.Atoi(splited[0])
	if err != nil {
		return err
	}
	if splited[1] != "Бег" && splited[1] != "Ходьба" {
		return errors.New("invalid type of training")
	}
	duration, err := time.ParseDuration(splited[2])
	if err != nil {
		return err
	}
	t.Steps = steps
	t.TrainingType = splited[1]
	t.Duration = duration
	return nil

}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps)
	if distance <= 0 {
		return "", errors.New("distance less or equal 0")
	}
	if t.TrainingType == "Бег" {
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", t.TrainingType, t.Duration.Hours(), spentenergy.Distance(t.Steps), spentenergy.MeanSpeed(t.Steps, t.Duration), spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)), nil
	} else if t.TrainingType == "Ходьба" {
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", t.TrainingType, t.Duration.Hours(), spentenergy.Distance(t.Steps), spentenergy.MeanSpeed(t.Steps, t.Duration), spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)), nil
	}
	return "неизвестный тип тренировки", errors.New("unknown training type")
}
