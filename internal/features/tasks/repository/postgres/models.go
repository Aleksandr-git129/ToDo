package task_postgres_repository

import (
	"time"
	"todo-list/internal/core/domain"
)

type TaskModel struct {
	ID           int
	Version      int
	Title        string
	Description  *string
	Completed    bool
	CreatedAt    time.Time
	CompletedAt  *time.Time
	AuthorUserID int
}

func TaskDomainFromModel(taskModel TaskModel) domain.Task {
	return domain.NewTask(
		taskModel.ID,
		taskModel.Version,
		taskModel.Title,
		taskModel.Description,
		taskModel.Completed,
		taskModel.CreatedAt,
		taskModel.CompletedAt,
		taskModel.AuthorUserID,
	)
}

func TasksDomainsFromModels(tasksModels []TaskModel) []domain.Task {
	domains := make([]domain.Task, len(tasksModels))

	for i, model := range tasksModels {
		domains[i] = TaskDomainFromModel(model)
	}

	return domains
}
