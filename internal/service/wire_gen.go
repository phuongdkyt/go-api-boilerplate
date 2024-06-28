// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"github.com/anhnmt/go-api-boilerplate/internal/infrastructure/gormgen"
	"github.com/anhnmt/go-api-boilerplate/internal/pkg/config"
	"github.com/anhnmt/go-api-boilerplate/internal/service/auth/business"
	"github.com/anhnmt/go-api-boilerplate/internal/service/auth/transport/grpc"
	"github.com/anhnmt/go-api-boilerplate/internal/service/session/repository/postgres/command"
	"github.com/anhnmt/go-api-boilerplate/internal/service/user/business"
	"github.com/anhnmt/go-api-boilerplate/internal/service/user/repository/postgres/command"
	"github.com/anhnmt/go-api-boilerplate/internal/service/user/repository/postgres/query"
	"github.com/anhnmt/go-api-boilerplate/internal/service/user/transport/grpc"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

// Injectors from wire.go:

func New(grpcSrv *grpc.Server, gormQuery *gormgen.Query, rdb redis.UniversalClient, cfg config.JWT) error {
	command := usercommand.New(gormQuery)
	query := userquery.New(gormQuery)
	business := userbusiness.New(command, query)
	userServiceServer := usergrpc.New(grpcSrv, business)
	sessioncommandCommand := sessioncommand.New(gormQuery)
	authbusinessBusiness := authbusiness.New(cfg, query, sessioncommandCommand)
	authServiceServer := authgrpc.New(grpcSrv, authbusinessBusiness)
	error2 := initServices(userServiceServer, authServiceServer)
	return error2
}
