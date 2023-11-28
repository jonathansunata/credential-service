// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import (
	"context"
	"github.com/SawitProRecruitment/UserService/domain/entity"
)

type RepositoryInterface interface {
	Update(ctx context.Context, input *entity.User) (err error)
	Register(ctx context.Context, input *entity.User) (id int32, err error)
	GetById(ctx context.Context, id int32) (user *entity.User, err error)
	GetByPhone(ctx context.Context, phone string) (user *entity.User, err error)
	GetByPhoneOther(ctx context.Context, phone string, id int32) (user *entity.User, err error)
}
