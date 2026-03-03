package postgres

import (
	"context"
	"database/sql"
	"errors"
	"librigo/internal/domain/user"
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
			INSERT INTO users (id, name, email,password_hash,role) 
			VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.ExecContext(ctx, query, user.ID, user.Name, user.Email, user.PasswordHash, user.Role)
	if err != nil {
		// PostgreSQLの固有エラーかどうかをチェック
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			// "23505" は unique_violation (一意制約違反) のコードです
			if pgErr.Code == "23505" {
				// userdomainにErrDuplicateUserを定義すべき
				return userdomain.ErrDuplicateUser
			}
		}
		// それ以外のエラー（接続不良など）はそのまま返す
		return err
	}
	return nil
}

// FindByEmailはメールアドレスでユーザーを検索するメソッドです。ユーザーが見つからない場合はuser.ErrUserNotFoundを返します。
func (r *UserRepository) FindByEmail(ctx context.Context, email user.Email) (*userdomain.User, error) {
	query := `SELECT id, name, email, password_hash, role FROM users WHERE email = $1`
	row := r.db.QueryRowContext(ctx, query, email)

	var u user.User
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, userdomain.ErrUserNotFound
		}
		return nil, err
	}
	return &u, nil
}

// FindByIDはユーザーIDでユーザーを検索するメソッドです。ユーザーが見つからない場合はuser.ErrUserNotFoundを返します。
func (r *UserRepository) FindByID(ctx context.Context, id userdomain.UserID) (*userdomain.User, error) {
	query := `SELECT id, name, email, password_hash, role FROM users WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var u user.User
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, userdomain.ErrUserNotFound
		}
		return nil, err
	}
	return &u, nil
}
