package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	core_logger "todo-list/internal/core/logger"
	core_pgx_pool "todo-list/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "todo-list/internal/core/transport/http/middleware"
	core_http_server "todo-list/internal/core/transport/http/server"
	statistics_postgres_repository "todo-list/internal/features/statistics/repository"
	statistics_service "todo-list/internal/features/statistics/service"
	statistics_transport_http "todo-list/internal/features/statistics/transport/http"
	task_postgres_repository "todo-list/internal/features/tasks/repository/postgres"
	task_service "todo-list/internal/features/tasks/service"
	tasks_transport "todo-list/internal/features/tasks/transport/http"
	users_postgres_repository "todo-list/internal/features/users/repository/postgres"
	users_service "todo-list/internal/features/users/service"
	users_transport_http "todo-list/internal/features/users/transport/http"

	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init application logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("initializing postgres pool connection")
	pool, err := core_pgx_pool.NewPool(ctx, core_pgx_pool.NewConfigMust())

	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initializing feature", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(usersService)

	logger.Debug("initializing feature", zap.String("feature", "tasks"))
	taskRepository := task_postgres_repository.NewTasksRepository(pool)
	taskService := task_service.NewTaskService(taskRepository)
	taskTranspornHTTP := tasks_transport.NewTasksHTTPHandler(taskService)

	logger.Debug("initializing feature", zap.String("feature", "statistics"))
	statisticsRepository := statistics_postgres_repository.NewStatisticsRepository(pool)
	statisticsService := statistics_service.NewStatisticsService(statisticsRepository)
	statisticsTransportHTTP := statistics_transport_http.NewStatisticsHTTPHandler(statisticsService)

	logger.Debug("initializing HTTP server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)
	apiVersionRouter.RegisterRoutes(taskTranspornHTTP.Routes()...)
	apiVersionRouter.RegisterRoutes(statisticsTransportHTTP.Routes()...)

	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
