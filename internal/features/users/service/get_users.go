package users_service

import (
	"context"
	"fmt"
	"todo-list/internal/core/domain"
	core_errors "todo-list/internal/core/errors"
)

func (s *UsersService) GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error) {
	if err := validateQueryParams(limit, offset); err != nil {
		return nil, fmt.Errorf("failed to validate query params", err)
	}

	users, err := s.usersRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get users from repository: %w", err)
	}

	return users, nil
}

func validateQueryParams(limit *int, offset *int) error {
	if limit != nil && *limit < 0 {
		return fmt.Errorf("limit must be non-negative: %w", core_errors.ErrInvalidArgument)
	}

	if offset != nil && *offset < 0 {
		return fmt.Errorf("offset must be non-negative: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}
