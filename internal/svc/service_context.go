package svc

import (
	"muxi-empolyment/internal/config"
	"muxi-empolyment/internal/middleware"
	"muxi-empolyment/internal/pkg/ijwt"
)

type ServiceContext struct {
	Config config.Config
	JWTHandler ijwt.JWTHandler
	AuthMiddleware middleware.AuthMiddleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	JWTHandler:=ijwt.NewJWTHandler(c.Auth.AccessSecret)
	AuthMiddleware:=middleware.NewAuthMiddleware(c,JWTHandler)
	return &ServiceContext{
		Config: c,
		JWTHandler: JWTHandler,
		AuthMiddleware: AuthMiddleware,
	}
}
