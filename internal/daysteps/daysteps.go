package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/s444v/go-fifth-sprint/internal/personaldata"
	"github.com/s444v/go-fifth-sprint/internal/spentenergy"
)

const (
	StepLength = 0.65
)

// создайте структуру DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {
	// ваш код ниже
	splited := strings.Split(datastring, ",")
	if len(splited) != 2 {
		return errors.New("string format does not match")
	}
	steps, err := strconv.Atoi(splited[0])
	if steps <= 0 {
		return errors.New("steps less or equal 0")
	}
	if err != nil {
		return err
	}
	duration, err := time.ParseDuration(splited[1])
	if err != nil {
		return err
	}
	ds.Steps = steps
	ds.Duration = duration
	return nil
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", errors.New("duration less or equal 0")
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", ds.Steps, spentenergy.Distance(ds.Steps), spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)), nil
}
