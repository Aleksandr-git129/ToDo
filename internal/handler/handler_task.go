package handler

import (
	"net/http"
	"todo-list/internal/model"
)

type HandlerTask struct {
	Task *model.Task
}

func HandlerTaskConsctructor(t *model.Task) *HandlerTask {
	return &HandlerTask{
		Task: t,
	}
}

func (h *HandlerTask) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *HandlerTask) GetTaskHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *HandlerTask) GetAllTaskHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *HandlerTask) GetUserTaskHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *HandlerTask) ChangeTaskHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *HandlerTask) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {

}
