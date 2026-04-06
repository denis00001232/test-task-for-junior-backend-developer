package handlers

import (
	"time"

	taskdomain "example.com/taskservice/internal/domain/task"
)

type taskMutationDTO struct {
	Title           string                     `json:"title"`
	Description     string                     `json:"description"`
	Status          taskdomain.Status          `json:"status"`
	PeriodicityType taskdomain.PeriodicityType `json:"periodicity_type"`
	DailyInterval   int                        `json:"daily_interval,omitempty"`
	MonthlyDays     []int                      `json:"monthly_days,omitempty"`
	SpecificDates   []time.Time                `json:"specific_dates,omitempty"`
	OddEven         taskdomain.OddEvenType     `json:"odd_even_type,omitempty"`
}

type taskDTO struct {
	ID              int64                      `json:"id"`
	Title           string                     `json:"title"`
	Description     string                     `json:"description"`
	Status          taskdomain.Status          `json:"status"`
	CreatedAt       time.Time                  `json:"created_at"`
	UpdatedAt       time.Time                  `json:"updated_at"`
	PeriodicityType taskdomain.PeriodicityType `json:"periodicity_type"`
	DailyInterval   int                        `json:"daily_interval,omitempty"`
	MonthlyDays     []int                      `json:"monthly_days,omitempty"`
	SpecificDates   []time.Time                `json:"specific_dates,omitempty"`
	OddEven         taskdomain.OddEvenType     `json:"odd_even_type,omitempty"`
}

func newTaskDTO(task *taskdomain.Task) taskDTO {
	return taskDTO{
		ID:              task.ID,
		Title:           task.Title,
		Description:     task.Description,
		Status:          task.Status,
		CreatedAt:       task.CreatedAt,
		UpdatedAt:       task.UpdatedAt,
		PeriodicityType: task.PeriodicityType,
		DailyInterval:   task.DailyInterval,
		MonthlyDays:     task.MonthlyDays,
		SpecificDates:   task.SpecificDates,
		OddEven:         task.OddEven,
	}
}
