package task

import "time"

type Status string

const (
	StatusNew        Status = "new"
	StatusInProgress Status = "in_progress"
	StatusDone       Status = "done"
)

type OddEvenType string

const (
	Odd  OddEvenType = "odd"
	Even OddEvenType = "even"
)

type PeriodicityType string

const (
	DailyInterval PeriodicityType = "daily interval"
	MonthlyDays   PeriodicityType = "monthly days"
	SpecificDates PeriodicityType = "specific dates"
	OddEven       PeriodicityType = "odd even"
)

type Task struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	PeriodicityType PeriodicityType `json:"periodicity_type"`
	DailyInterval   int             `json:"daily_interval,omitempty"` // каждый N-й день
	MonthlyDays     []int           `json:"monthly_days,omitempty"`   // дни месяца [1, 15, 30]
	SpecificDates   []time.Time     `json:"specific_dates,omitempty"` // конкретные даты ["2024-01-01"]
	OddEven         OddEvenType     `json:"odd_even_type,omitempty"`  // "odd" или "even"
}

func (s Status) Valid() bool {
	switch s {
	case StatusNew, StatusInProgress, StatusDone:
		return true
	default:
		return false
	}
}

func (pt PeriodicityType) Valid() bool {
	switch pt {
	case DailyInterval, MonthlyDays, SpecificDates, OddEven:
		return true
	default:
		return false
	}
}

func (oet OddEvenType) Valid() bool {
	switch oet {
	case Odd, Even:
		return true
	default:
		return false
	}
}
