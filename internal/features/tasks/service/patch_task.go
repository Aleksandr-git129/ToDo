package task_service

import (
	"context"
	"fmt"
	"todo-list/internal/core/domain"
)

func (s *TaskService) PatchTask(ctx context.Context, id int, taskPatch domain.TaskPatch) (domain.Task, error) {
	task, err := s.taskRepository.GetTask(ctx, id)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task: %w", err)
	}

	if err := task.ApplyPatch(taskPatch); err != nil {
		return domain.Task{}, fmt.Errorf("task applyPatch: %w", err)
	}

	patchedTask, err := s.taskRepository.PatchTask(ctx, id, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("patch task: %w", err)
	}

	return patchedTask, nil
}