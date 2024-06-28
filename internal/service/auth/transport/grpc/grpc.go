package authgrpc

import (
	"context"

	"google.golang.org/grpc"

	authbusiness "github.com/anhnmt/go-api-boilerplate/internal/service/auth/business"
	"github.com/anhnmt/go-api-boilerplate/proto/pb"
)

type grpcService struct {
	pb.UnimplementedAuthServiceServer

	authBusiness *authbusiness.Business
}

func New(
	grpcSrv *grpc.Server,
	authBusiness *authbusiness.Business,
) pb.AuthServiceServer {
	svc := &grpcService{
		authBusiness: authBusiness,
	}

	pb.RegisterAuthServiceServer(grpcSrv, svc)

	return svc
}

func (s *grpcService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return s.authBusiness.Login(ctx, req)
}

func (s *grpcService) Info(ctx context.Context, _ *pb.InfoRequest) (*pb.InfoResponse, error) {
	return s.authBusiness.Info(ctx)
}

func (s *grpcService) RevokeToken(ctx context.Context, _ *pb.RevokeTokenRequest) (*pb.RevokeTokenResponse, error) {
	return nil, s.authBusiness.RevokeToken(ctx)
}

func (s *grpcService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	return s.authBusiness.RefreshToken(ctx, req)
}
