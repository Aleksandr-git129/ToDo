package statistics_transport_http

import (
	"context"
	"net/http"
	"time"
	"todo-list/internal/core/domain"
	core_http_server "todo-list/internal/core/transport/http/server"
)

type StatitsticsHTTPHandler struct {
	statisticsSevice StatisticsSevice
}

type StatisticsSevice interface {
	GetStatistics(
		ctx context.Context,
		userID *int,
		form *time.Time,
		to *time.Time,
	) (domain.Statistics, error)
}

func NewStatisticsHTTPHandler(statisticsSevice StatisticsSevice) *StatitsticsHTTPHandler {
	return &StatitsticsHTTPHandler{
		statisticsSevice: statisticsSevice,
	}
}

func (h *StatitsticsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/statistics",
			Handler: h.GetStatistics,
		},
	}
}
