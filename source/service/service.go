package service

import (
	"basic/pkg/jwt"
	logger "basic/pkg/logger"

	"github.com/spf13/viper"
)

type Service struct {
	logger *logger.Logger
	jwt    *jwt.JWT
	conf   *viper.Viper
}

func NewService(logger *logger.Logger, jwt *jwt.JWT, conf *viper.Viper) *Service {
	return &Service{
		logger: logger,
		jwt:    jwt,
		conf:   conf,
	}
}
