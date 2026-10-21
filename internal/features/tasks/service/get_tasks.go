package task_service

import (
	"context"
	"fmt"
	"todo-list/internal/core/domain"
	core_errors "todo-list/internal/core/errors"
)

func (s *TaskService) GetTasks(ctx context.Context, userID *int, limit *int, offset *int) ([]domain.Task, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf("limit must be non-negative: %w", core_errors.ErrInvalidArgument)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf("offset must be non-negative: %w", core_errors.ErrInvalidArgument)
	}

	tasks, err := s.taskRepository.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get task from repository; %w", err)
	}

	return tasks, nil
}
