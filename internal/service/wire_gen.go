// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"github.com/anhnmt/go-api-boilerplate/internal/infrastructure/gormgen"
	"github.com/anhnmt/go-api-boilerplate/internal/service/user/repository/postgres/command"
	"github.com/anhnmt/go-api-boilerplate/internal/service/user/transport/grpc"
	"google.golang.org/grpc"
)

// Injectors from wire.go:

func New(grpcSrv *grpc.Server, gormQuery *gormgen.Query) error {
	command := usercommand.New(gormQuery)
	userServiceServer := usergrpc.New(grpcSrv, command)
	error2 := initServices(userServiceServer)
	return error2
}
