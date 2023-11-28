package repository

import (
	"context"
	"fmt"
	"github.com/SawitProRecruitment/UserService/domain/entity"
	error_handler "github.com/SawitProRecruitment/UserService/error"
)

func (r *Repository) Register(ctx context.Context, input *entity.User) (int32, error) {
	id := int32(0)
	err := r.Db.QueryRowContext(ctx, "INSERT INTO users (full_name, phone_number, password, salt) VALUES ($1, $2, $3, $4) RETURNING id",
		input.FullName, input.PhoneNumber, input.Password, input.Salt).Scan(&id)
	if err != nil {
		fmt.Errorf("error when register user with id : %v, err: %v", input.Id, err)
		return 0, error_handler.NewCustomError(500, err.Error())
	}
	return id, nil
}

func (r *Repository) Update(ctx context.Context, input *entity.User) error {
	_, err := r.Db.ExecContext(ctx, "UPDATE users SET full_name = $1, phone_number = $2, password = $3, salt = $4, count_login = $5 WHERE id = $6",
		input.FullName, input.PhoneNumber, input.Password, input.Salt, input.CountLogin, input.Id)
	if err != nil {
		fmt.Errorf("error when get user by id : %v, err: %v", input.FullName, err)
		return error_handler.NewCustomError(500, err.Error())
	}
	return nil
}

func (r *Repository) GetById(ctx context.Context, id int32) (user *entity.User, err error) {
	user = &entity.User{}
	err = r.Db.QueryRowContext(ctx, "SELECT id, full_name, phone_number, password, salt FROM users WHERE id = $1", id).Scan(&user.Id,
		&user.FullName, &user.PhoneNumber, &user.Password, &user.Salt)
	if err != nil {
		fmt.Errorf("error when get user by id : %v, err: %v", id, err)
		return nil, error_handler.NewCustomError(500, err.Error())
	}
	return user, nil
}

func (r *Repository) GetByPhone(ctx context.Context, phone string) (*entity.User, error) {
	user := &entity.User{}
	err := r.Db.QueryRowContext(ctx, "SELECT id, full_name, phone_number, password, salt, count_login FROM users WHERE phone_number = $1", phone).Scan(&user.Id,
		&user.FullName, &user.PhoneNumber, &user.Password, &user.Salt, &user.CountLogin)
	if err != nil {
		fmt.Errorf("error when get user by phone with id : %v, err: %v", phone, err)
		return nil, error_handler.NewCustomError(500, err.Error())
	}
	return user, nil
}

func (r *Repository) GetByPhoneOther(ctx context.Context, phone string, id int32) (*entity.User, error) {
	user := &entity.User{}
	err := r.Db.QueryRowContext(ctx, "SELECT id, full_name, phone_number, password, salt FROM users WHERE phone_number = $1 AND id != $2", phone, id).Scan(&user.Id,
		&user.FullName, &user.PhoneNumber, &user.Password, &user.Salt)
	if err != nil {
		fmt.Errorf("error when get user by phone other with id : %v, err: %v", phone, err)
		return nil, error_handler.NewCustomError(500, err.Error())
	}
	return user, nil
}
