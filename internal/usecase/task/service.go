package task

import (
	"context"
	"fmt"
	"strings"
	"time"

	taskdomain "example.com/taskservice/internal/domain/task"
)

type Service struct {
	repo Repository
	now  func() time.Time
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
		now:  func() time.Time { return time.Now().UTC() },
	}
}

func (s *Service) Create(ctx context.Context, input CreateInput) (*taskdomain.Task, error) {
	normalized, err := validateCreateInput(input)
	if err != nil {
		return nil, err
	}

	model := &taskdomain.Task{
		Title:           normalized.Title,
		Description:     normalized.Description,
		Status:          normalized.Status,
		PeriodicityType: normalized.PeriodicityType,
		DailyInterval:   normalized.DailyInterval,
		MonthlyDays:     normalized.MonthlyDays,
		SpecificDates:   normalized.SpecificDates,
		OddEven:         normalized.OddEven,
	}

	now := s.now()
	model.CreatedAt = now
	model.UpdatedAt = now

	created, err := s.repo.Create(ctx, model)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (s *Service) GetByID(ctx context.Context, id int64) (*taskdomain.Task, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: id must be positive", ErrInvalidInput)
	}

	return s.repo.GetByID(ctx, id)
}

func (s *Service) Update(ctx context.Context, id int64, input UpdateInput) (*taskdomain.Task, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: id must be positive", ErrInvalidInput)
	}

	normalized, err := validateUpdateInput(input)
	if err != nil {
		return nil, err
	}

	model := &taskdomain.Task{
		ID:              id,
		Title:           normalized.Title,
		Description:     normalized.Description,
		Status:          normalized.Status,
		UpdatedAt:       s.now(),
		PeriodicityType: normalized.PeriodicityType,
		DailyInterval:   normalized.DailyInterval,
		MonthlyDays:     normalized.MonthlyDays,
		SpecificDates:   normalized.SpecificDates,
		OddEven:         normalized.OddEven,
	}

	updated, err := s.repo.Update(ctx, model)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("%w: id must be positive", ErrInvalidInput)
	}

	return s.repo.Delete(ctx, id)
}

func (s *Service) List(ctx context.Context) ([]taskdomain.Task, error) {
	return s.repo.List(ctx)
}

func validateCreateInput(input CreateInput) (CreateInput, error) {
	input.Title = strings.TrimSpace(input.Title)
	input.Description = strings.TrimSpace(input.Description)

	if input.Title == "" {
		return CreateInput{}, fmt.Errorf("%w: title is required", ErrInvalidInput)
	}

	if input.Status == "" {
		input.Status = taskdomain.StatusNew
	}

	if !input.Status.Valid() {
		return CreateInput{}, fmt.Errorf("%w: invalid status", ErrInvalidInput)
	}

	if !input.PeriodicityType.Valid() {
		return CreateInput{}, fmt.Errorf("%w: invalid periodicity type", ErrInvalidInput)
	}

	if err := validatePeriodicityType(input.PeriodicityType,
		input.DailyInterval,
		input.MonthlyDays,
		input.SpecificDates,
		input.OddEven); err != nil {
		return CreateInput{}, err
	}

	return input, nil
}

func validateUpdateInput(input UpdateInput) (UpdateInput, error) {
	input.Title = strings.TrimSpace(input.Title)
	input.Description = strings.TrimSpace(input.Description)

	if input.Title == "" {
		return UpdateInput{}, fmt.Errorf("%w: title is required", ErrInvalidInput)
	}

	if !input.Status.Valid() {
		return UpdateInput{}, fmt.Errorf("%w: invalid status", ErrInvalidInput)
	}

	if err := validatePeriodicityType(input.PeriodicityType,
		input.DailyInterval,
		input.MonthlyDays,
		input.SpecificDates,
		input.OddEven); err != nil {
		return UpdateInput{}, err
	}

	return input, nil
}

func validatePeriodicityType(periodicityType taskdomain.PeriodicityType, dailyInterval int, monthlyDays []int, specificDates []time.Time, oddEven taskdomain.OddEvenType) error {
	if !periodicityType.Valid() {
		return fmt.Errorf("%w: invalid periodicity type", ErrInvalidInput)
	}

	switch periodicityType {
	case taskdomain.DailyInterval:
		if dailyInterval < 0 {
			return fmt.Errorf("%w: daily interval must be positive", ErrInvalidInput)
		}

	case taskdomain.MonthlyDays:
		if len(monthlyDays) == 0 {
			return fmt.Errorf("%w: monthly days is required", ErrInvalidInput)
		}
		for _, num := range monthlyDays {
			if num < 1 || num > 31 {
				return fmt.Errorf("%w: incorrect monthly day", ErrInvalidInput)
			}
		}

	case taskdomain.SpecificDates:
		if len(specificDates) == 0 {
			return fmt.Errorf("%w: specific days is required", ErrInvalidInput)
		}

	case taskdomain.OddEven:
		if !oddEven.Valid() {
			return fmt.Errorf("%w: invalid odd even", ErrInvalidInput)
		}
	}

	return nil
}
