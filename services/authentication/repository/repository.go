package repository

import (
	"context"
	"database/sql"
	"github.com/ungame/go-signup/db"
	"github.com/ungame/go-signup/services/authentication/entities"
	"log"
	"time"
)

type AuthenticationUsersRepository interface {
	Create(ctx context.Context, entity *entities.DbAuthenticationUser) error
	Update(ctx context.Context, entity *entities.DbAuthenticationUser) error
	Get(ctx context.Context, column, value string) (*entities.DbAuthenticationUser, error)
	GetAll(ctx context.Context, limit, offset int) (*entities.DbAuthenticationUser, error)
	Delete(ctx context.Context, id string) error
}

type authenticationUsersRepository struct {
	conn                          db.Connection
	createAuthenticationUsersStmt *sql.Stmt
	updateAuthenticationUsersStmt *sql.Stmt
	deleteAuthenticationUsersStmt *sql.Stmt
}

func NewAuthenticationUsersRepository(conn db.Connection) AuthenticationUsersRepository {
	r := &authenticationUsersRepository{conn: conn}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var err error

	r.createAuthenticationUsersStmt, err = r.conn.PrepareContext(ctx, createAuthenticationUsersQuery)
	if err != nil {
		log.Fatalln(err)
	}

	r.updateAuthenticationUsersStmt, err = r.conn.PrepareContext(ctx, updateAuthenticationUsersQuery)
	if err != nil {
		log.Fatalln(err)
	}

	r.deleteAuthenticationUsersStmt, err = r.conn.PrepareContext(ctx, deleteAuthenticationUsersQuery)
	if err != nil {
		log.Fatalln(err)
	}

	return r
}

func (r *authenticationUsersRepository) Close() {
	if err := r.createAuthenticationUsersStmt.Close(); err != nil {
		log.Println("error on close createAuthenticationUsersStmt:", err)
	}
	if err := r.updateAuthenticationUsersStmt.Close(); err != nil {
		log.Println("error on close updateAuthenticationUsersStmt:", err)
	}
	if err := r.deleteAuthenticationUsersStmt.Close(); err != nil {
		log.Println("error on close deleteAuthenticationUsersStmt:", err)
	}
}

func (r *authenticationUsersRepository) Create(ctx context.Context, entity *entities.DbAuthenticationUser) error {
	_, err := r.createAuthenticationUsersStmt.ExecContext(
		ctx,
		entity.Id,
		entity.Email,
		entity.Username,
		entity.Phone,
		entity.Password,
		entity.CreatedAt,
		entity.UpdatedAt,
	)
	return err
}

func (r *authenticationUsersRepository) Update(ctx context.Context, entity *entities.DbAuthenticationUser) error {
	_, err := r.updateAuthenticationUsersStmt.ExecContext(
		ctx,
		entity.Email,
		entity.Username,
		entity.Password,
		entity.Phone,
		entity.UpdatedAt,
	)
	return err
}

func (r *authenticationUsersRepository) Get(ctx context.Context, column, value string) (*entities.DbAuthenticationUser, error) {
	row := r.conn.QueryRowContext(ctx, getAuthenticationUserQuery, column, value)

	var entity entities.DbAuthenticationUser

	err := row.Scan(
		&entity.Id,
		&entity.Email,
		&entity.Username,
		&entity.Password,
		&entity.Phone,
		&entity.CreatedAt,
		&entity.UpdatedAt,
	)

	return &entity, err
}

func (r *authenticationUsersRepository) GetAll(ctx context.Context, limit, offset int) (*entities.DbAuthenticationUser, error) {
	row, err := r.conn.QueryContext(ctx, getAllAuthenticationUsersQuery, limit, offset)
	if err != nil {
		return nil, err
	}

	var entity entities.DbAuthenticationUser

	err = row.Scan(
		&entity.Id,
		&entity.Email,
		&entity.Username,
		&entity.Password,
		&entity.Phone,
		&entity.CreatedAt,
		&entity.UpdatedAt,
	)

	return &entity, err
}

func (r *authenticationUsersRepository) Delete(ctx context.Context, id string) error {
	_, err := r.deleteAuthenticationUsersStmt.ExecContext(ctx, id)
	return err
}
