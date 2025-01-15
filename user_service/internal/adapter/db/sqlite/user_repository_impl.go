package sqlite

import (
	"context"

	"github.com/rcarvalho-pb/mottu-user_service/internal/model"
)

func (db *DB) CreateUser(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `INSERT INTO tb_users
    (username, password, name, birth_date, cnpj, cnh, cnh_type, cnh_file_path)
VALUES
	(:username, :password, :name, :birth_date, :cnpj, :cnh, :cnh_type, :cnh_file_path)`
	if _, err := db.DB.NamedExecContext(ctx, stmt, &user); err != nil {
		return err
	}
	return nil
}

func (db *DB) UpdateUser(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `UPDATE tb_users SET
	username = :username, password = :password, name = :name, role := role, birth_date = :birth_date,
	cnpj = :cnpj, cnh = :cnh, cnh_type = :cnh_type, cnh_file_path = :cnh_file_path, active_location = :active_location,
	updated_at = :updated_at, active = :active
	WHERE id = :id`
	if _, err := db.DB.NamedExecContext(ctx, stmt, &user); err != nil {
		return err
	}
	return nil
}

func (db *DB) GetUserById(id int64) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `SELECT * FROM tb_users WHERE id = ?`
	var user *model.User
	if err := db.DB.GetContext(ctx, &user, stmt, id); err != nil {
		return nil, err
	}
	return user, nil
}

func (db *DB) GetUserByUsername(username string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `SELECT * FROM tb_users WHERE username = ?`
	var user *model.User
	if err := db.DB.GetContext(ctx, &user, stmt, username); err != nil {
		return nil, err
	}
	return user, nil
}

func (db *DB) GetAllUsers() ([]*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `SELECT * FROM tb_users`
	var users []*model.User
	if err := db.DB.SelectContext(ctx, &users, stmt); err != nil {
		return nil, err
	}
	return users, nil
}
