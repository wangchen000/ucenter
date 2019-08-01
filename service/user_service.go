package service

import (
	"context"
	"ucenter/rpc"
)

type UserService struct{}

func (u *UserService) Registry(cxt context.Context,registry *rpc.Register) (*rpc.UserDto, error) {
	return &rpc.UserDto{}, nil
}