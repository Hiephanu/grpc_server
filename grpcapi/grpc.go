package api

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcApi struct {
	UnimplementedUserServiceServer
}

func (s *GrpcApi) GetAllUsers(ctx context.Context, in *emptypb.Empty) (*UserListResponse, error) {
	// Cài đặt logic để lấy tất cả người dùng
	return &UserListResponse{}, nil
}

func (s *GrpcApi) PostUser(ctx context.Context, in *UserPostRequest) (*UserPostResponse, error) {
	// Cài đặt logic để thêm người dùng
	return &UserPostResponse{}, nil
}

func (s *GrpcApi) UpdateUser(ctx context.Context, in *UserUpdateRequest) (*UserResponse, error) {
	// Cài đặt logic để cập nhật người dùng
	return &UserResponse{}, nil
}

func (s *GrpcApi) DeleteUser(ctx context.Context, in *UserDeleteRequest) (*UserDeleteResponse, error) {
	// Cài đặt logic để xóa người dùng
	return &UserDeleteResponse{}, nil
}
