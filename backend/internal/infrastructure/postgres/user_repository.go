package postgres

import (
	"context"
	"database/sql"
	"errors"
	userdomain "librigo/internal/domain/user"

	"github.com/jackc/pgx/v5/pgconn"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) userdomain.UserRepository {
	return &UserRepository{db: db}
}

// Saveはユーザーをデータベースに保存するメソッドです。ユーザーのメールアドレスが既に存在する場合はuser.ErrDuplicateUserを返します。
func (r *UserRepository) Save(ctx context.Context, user *userdomain.User) error {
	query := `
			INSERT INTO users (id, name, email, password_hash, role) 
			VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.ExecContext(ctx, query,
		user.ID().String(),
		user.Name().String(),
		user.Email().String(),
		user.PasswordHash().String(),
		user.Role().String(),
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return userdomain.ErrDuplicateUser
			}
		}
		return err
	}
	return nil
}

// FindByEmailはメールアドレスでユーザーを検索するメソッドです。ユーザーが見つからない場合はuser.ErrUserNotFoundを返します。
func (r *UserRepository) FindByEmail(ctx context.Context, email userdomain.UserEmail) (*userdomain.User, error) {
	query := `SELECT id, name, email, password_hash, role FROM users WHERE email = $1`
	row := r.db.QueryRowContext(ctx, query, email.String())

	var (
		resID           string
		resName         string
		resEmail        string
		resPasswordHash string
		resRole         string
	)

	err := row.Scan(&resID, &resName, &resEmail, &resPasswordHash, &resRole)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, userdomain.ErrUserNotFound
		}
		return nil, err
	}

	return userdomain.NewUser(resID, resName, resEmail, resPasswordHash, resRole)
}

// FindByIDはユーザーIDでユーザーを検索するメソッドです。ユーザーが見つからない場合はuser.ErrUserNotFoundを返します。
func (r *UserRepository) FindByID(ctx context.Context, id userdomain.UserID) (*userdomain.User, error) {
	query := `SELECT id, name, email, password_hash, role FROM users WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id.String())

	var (
		resID           string
		resName         string
		resEmail        string
		resPasswordHash string
		resRole         string
	)

	err := row.Scan(&resID, &resName, &resEmail, &resPasswordHash, &resRole)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, userdomain.ErrUserNotFound
		}
		return nil, err
	}

	return userdomain.NewUser(resID, resName, resEmail, resPasswordHash, resRole)
}
