package postgresql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/SawitProRecruitment/UserService/model"
	"github.com/google/uuid"
)

// GetUserById implements repository.PostgresqlRepositoryInterface.
func (p *postgresqlRepository) GetUserById(ctx context.Context, id uuid.UUID) (model.User, error) {
	var user model.User
	if err := p.DB.QueryRowContext(ctx, "SELECT id, fullname, phonenumber FROM users WHERE id = $1", id).Scan(
		&user.Id,
		&user.Fullname,
		&user.Phonenumber,
	); err != nil {
		return user, fmt.Errorf("[postgresql][getuserbyid] error query: %w", err)
	}

	return user, nil
}

// GetUserByPhonenumber implements repository.PostgresqlRepositoryInterface.
func (p *postgresqlRepository) GetUserByPhonenumber(ctx context.Context, phonenumber string) (model.User, error) {
	var user model.User
	if err := p.DB.QueryRowContext(ctx, "SELECT id, fullname, phonenumber FROM users WHERE phonenumber = $1", phonenumber).Scan(
		&user.Id,
		&user.Fullname,
		&user.Phonenumber,
	); err != nil {
		return user, fmt.Errorf("[postgresql][getuserbyphonenumber] error query: %w", err)
	}

	return user, nil
}

// GetUserPasswordById implements repository.PostgresqlRepositoryInterface.
func (p *postgresqlRepository) GetUserPasswordById(ctx context.Context, id uuid.UUID) (model.UserPassword, error) {
	var userPassword model.UserPassword
	var hashedPassword string
	if err := p.DB.QueryRowContext(ctx, "SELECT user_id, password FROM user_password WHERE user_id = $1", id).Scan(
		&userPassword.Id,
		&hashedPassword,
	); err != nil {
		return userPassword, fmt.Errorf("[postgresql][getuserpasswordbyid] error query: %w", err)
	}

	userPassword.SetHashedPassword(hashedPassword)

	return userPassword, nil
}

// CheckPhonenumberExists implements repository.PostgresqlRepositoryInterface.
func (p *postgresqlRepository) CheckPhonenumberExists(ctx context.Context, phonenumber string) (bool, error) {
	var exists bool
	if err := p.DB.QueryRowContext(ctx, "SELECT EXISTS (SELECT * FROM users WHERE phonenumber = $1", phonenumber).Scan(
		&exists,
	); err != nil {
		return true, fmt.Errorf("[postgresql][checkphonenumberexists] error query: %v", err)
	}

	return exists, nil
}

// CountLoginSession implements repository.PostgresqlRepositoryInterface.
func (p *postgresqlRepository) CountLoginSession(ctx context.Context, userId uuid.UUID) error {
	tx, err := p.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("[postgresql][countloginsession] begin transaction error: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "INSERT INTO user_login (user_id, date) VALUES ($1, CURRENT_DATE) ON CONFLICT (user_id, date) DO UPDATE SET count = user_login.count + 1 WHERE user_login.user_id = $1 AND user_login.date = CURRENT_DATE", userId)
	if err != nil {
		return fmt.Errorf("[postgresql][countloginsession] error execute: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("[postgresql][countloginsession] commit error: %w", err)
	}

	return nil
}

// InsertNewUser implements repository.PostgresqlRepositoryInterface.
func (p *postgresqlRepository) InsertNewUser(ctx context.Context, newUser model.User, userAuth model.UserPassword) (uuid.UUID, error) {
	tx, err := p.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return uuid.Nil, fmt.Errorf("[postgresql][insertnewuser] begin transaction error: %w", err)
	}
	defer tx.Rollback()

	var userId uuid.UUID
	if err := tx.QueryRowContext(ctx, "INSERT INTO users (id, fullname, phonenumber) VALUES ($1, $2, $3) RETURNING id, fullname, phonenumber", newUser.Id, newUser.Fullname, newUser.Phonenumber).Scan(
		&userId,
	); err != nil {
		return uuid.Nil, fmt.Errorf("[postgresql][insertnewuser] execution error: %w", err)
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO user_password (user_id, password) VALUES ($1, $2)", userId, userAuth.GetPassword())
	if err != nil {
		return uuid.Nil, fmt.Errorf("[postgresql][insertnewuser] password execution error: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return uuid.Nil, fmt.Errorf("[postgresql][insertnewuser] commit error: %w", err)
	}

	return userId, nil
}

// UpdateUser implements repository.PostgresqlRepositoryInterface.
func (p *postgresqlRepository) UpdateUser(ctx context.Context, userId uuid.UUID, fullname string, phonenumber string) (uuid.UUID, error) {
	tx, err := p.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return uuid.Nil, fmt.Errorf("[postgresql][updateuser] begin transaction error: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "UPDATE users SET fullname = $2, phonenumber = $3 WHERE id = $1", userId, fullname, phonenumber)
	if err != nil {
		return uuid.Nil, fmt.Errorf("[postgresql][updateuser] execution error: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return uuid.Nil, fmt.Errorf("[postgresql][updateuser] commit error: %w", err)
	}

	return userId, nil
}
