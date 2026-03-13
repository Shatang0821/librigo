package user

import (
	"errors"
	"librigo/internal/domain/apperror"
	"net/mail"
	"strings"
	"unicode"
)

// ドメインエラー
var (
	ErrInvalidUserID       = apperror.New("INVALID_USER_ID", apperror.TypeInvalid)
	ErrInvalidUserName     = apperror.New("INVALID_USER_NAME", apperror.TypeInvalid)
	ErrInvalidUserEmail    = apperror.New("INVALID_USER_EMAIL", apperror.TypeInvalid)
	ErrWeakUserPassword    = apperror.New("WEAK_USER_PASSWORD", apperror.TypeInvalid)
	ErrInvalidUserPassword = apperror.New("INVALID_USER_PASSWORD", apperror.TypeInvalid)
	ErrInvalidUserRole     = apperror.New("INVALID_USER_ROLE", apperror.TypeInvalid)
)

// ユーザーの一意識別子
type UserID struct{ value string }

func NewUserID(v string) (UserID, error) {
	if v == "" {
		return UserID{}, ErrInvalidUserID.Wrap(errors.New("User ID can not be empty"))
	}
	return UserID{value: v}, nil
}

func (id UserID) String() string {
	return id.value
}

// ユーザーの名前
type UserName struct{ value string }

func NewUserName(v string) (UserName, error) {
	if v == "" {
		return UserName{}, ErrInvalidUserName.Wrap(errors.New("User name can not be empty"))
	}
	return UserName{value: v}, nil
}

func (n UserName) String() string {
	return n.value
}

// ユーザーのメールアドレス
type UserEmail struct{ value string }

func NewEmail(v string) (UserEmail, error) {
	trimmed := strings.TrimSpace(v)

	addr, err := mail.ParseAddress(trimmed)
	if err != nil {
		return UserEmail{}, ErrInvalidUserEmail.Wrap(errors.New("User email can not be empty"))
	}
	return UserEmail{value: addr.Address}, nil
}

func (e UserEmail) String() string {
	return e.value
}

// 生のパスワード
type UserRawPassword struct{ value string }

func NewUserRawPassword(v string) (UserRawPassword, error) {
	var (
		hasUpper bool
		hasLower bool
		hasDigit bool
	)

	// 長さチェック
	if len(v) < 8 {
		return UserRawPassword{}, ErrWeakUserPassword.Wrap(errors.New("password must be at least 8 characters"))
	}

	// 文字種チェック（一回のループで判定）
	for _, r := range v {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		}
	}

	// 3. 全ての条件を満たしているか判定
	// 要件に応じて「記号は任意」とするなら hasSpecial は外してもOKです
	if !hasUpper || !hasLower || !hasDigit {
		return UserRawPassword{}, ErrWeakUserPassword.Wrap(
			errors.New("password must contain uppercase, lowercase and digits"),
		)
	}
	return UserRawPassword{value: v}, nil
}

func (p UserRawPassword) String() string {
	return p.value
}

// ハッシュ化されたパスワード
type UserHashedPassword struct{ value string }

func NewUserHashedPassword(v string) (UserHashedPassword, error) {
	if v == "" {
		return UserHashedPassword{}, ErrInvalidUserPassword.Wrap(errors.New("User password can not be empty"))
	}
	return UserHashedPassword{value: v}, nil
}

func (p UserHashedPassword) String() string {
	return p.value
}

// UserRole はユーザーの権限を定義します
type UserRole struct{ value string }

// パッケージ外からはこれらの変数を通じてのみ権限を参照できる
var (
	RoleAdmin  = UserRole{value: "admin"}
	RoleMember = UserRole{value: "member"}
)

func NewUserRole(v string) (UserRole, error) {
	switch v {
	case "admin":
		return RoleAdmin, nil
	case "member":
		return RoleMember, nil
	default:
		return UserRole{}, ErrInvalidUserRole.Wrap(errors.New("Unknown user role: " + v))
	}
}

func (r UserRole) String() string {
	return r.value
}

func (r UserRole) IsAdmin() bool {
	return r.value == RoleAdmin.value
}
