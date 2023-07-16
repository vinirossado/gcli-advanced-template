package service

import (
	logger "basic/pkg/logger"
	"basic/source/middleware"
)

type Service struct {
	logger *logger.Logger
	jwt    *middleware.JWT
}

func NewService(logger *logger.Logger, jwt *middleware.JWT) *Service {
	return &Service{
		logger: logger,
		jwt:    jwt,
	}
}
