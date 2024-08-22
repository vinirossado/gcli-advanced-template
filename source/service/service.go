package service

import (
	"basic/pkg/jwt"
	logger "basic/pkg/logger"
)

type Service struct {
	logger *logger.Logger
	jwt    *jwt.JWT
}

func NewService(logger *logger.Logger, jwt *jwt.JWT) *Service {
	return &Service{
		logger: logger,
		jwt:    jwt,
	}
}
