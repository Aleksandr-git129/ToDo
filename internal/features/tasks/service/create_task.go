package task_service

import (
	"context"
	"fmt"
	"todo-list/internal/core/domain"
)

func (s *TaskService) CreateTask(
	ctx context.Context,
	task domain.Task,
) (domain.Task, error) {
	if err := task.Validate(); err !=nil {
		return domain.Task{}, fmt.Errorf("task domain validate: %w", err)
	}

	task, err := s.taskRepository.CreateTask(ctx, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("create task: %w", err)
	}

	return task, nil
}
