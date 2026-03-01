package usecase

import (
	"context"
	"errors"
	userdomain "librigo/internal/domain/user"
	"time"
)

// CreateUserInputはユーザー作成のための入力データを表す構造体です。
type SignUpInput struct {
	Name     string
	Email    string
	Password string
}

// CreateUserOutputはユーザー作成の結果を表す構造体です。
type SignUpOutput struct {
	ID    string
	Name  string
	Email string
}

type SignInInput struct {
	Email    string
	Password string
}

type SignInOutput struct {
	Token string // 後にJWTトークンを入れます
	User  *UserOutput
}

type UserOutput struct {
	ID    string
	Name  string
	Email string
	Role  userdomain.UserRole
}

type UserUseCase interface {
	SignUp(ctx context.Context, input SignUpInput) (*SignUpOutput, error)
	SignIn(ctx context.Context, input SignInInput) (*SignInOutput, error)
	GetUserByEmail(ctx context.Context, email string) (*UserOutput, error)
	GetUserByID(ctx context.Context, id string) (*UserOutput, error)
}

// userInteractorはUserUseCaseの実装です。
type userInteractor struct {
	repo           userdomain.UserRepository
	passwordHasher userdomain.PassWordHasher
	idGen          userdomain.IDGenerator
	tokenGen       userdomain.TokenGenerator
}

func NewUserUseCase(repo userdomain.UserRepository, passwordHasher userdomain.PassWordHasher, idGen userdomain.IDGenerator, tokenGen userdomain.TokenGenerator) UserUseCase {
	return &userInteractor{
		repo:           repo,
		passwordHasher: passwordHasher,
		idGen:          idGen,
		tokenGen:       tokenGen,
	}
}

// ユーザの新規登録
func (i *userInteractor) SignUp(ctx context.Context, in SignUpInput) (*SignUpOutput, error) {
	// パスワードの強度をチェック
	if err := userdomain.ValidatePassword(in.Password); err != nil {
		return nil, err
	}

	// emailの重複をチェック
	existingUser, _ := i.repo.FindByEmail(ctx, in.Email)
	if existingUser != nil {
		return nil, userdomain.ErrDuplicateUser
	}

	// パスワードをハッシュ化
	hashedPassword, err := i.passwordHasher.Hash(in.Password)
	if err != nil {
		return nil, err
	}

	newUser, err := userdomain.NewUser(
		i.idGen.Generate(),
		in.Name,
		in.Email,
		hashedPassword,
		userdomain.RoleMember, // デフォルトでMemberロールを設定
	)
	if err != nil {
		return nil, err
	}

	// ユーザーを保存
	if err := i.repo.Save(ctx, newUser); err != nil {
		return nil, err
	}

	return &SignUpOutput{
		ID:    string(newUser.ID),
		Name:  newUser.Name,
		Email: newUser.Email,
	}, nil
}

// ユーザのサインイン
func (i *userInteractor) SignIn(ctx context.Context, in SignInInput) (*SignInOutput, error) {
	// メールアドレスでユーザを取得
	user, err := i.repo.FindByEmail(ctx, in.Email)
	if err != nil {
		if errors.Is(err, userdomain.ErrUserNotFound) {
			return nil, userdomain.ErrInvalidCredentials
		}
		return nil, err
	}

	// パスワードを検証
	if err := i.passwordHasher.Compare(user.PasswordHash, in.Password); err != nil {
		return nil, userdomain.ErrInvalidCredentials
	}

	// JWTトークンを生成 有効期限は1時間
	token, err := i.tokenGen.Generate(user, 1*time.Hour)
	if err != nil {
		return nil, err
	}

	return &SignInOutput{
		Token: token,
		User: &UserOutput{
			ID:    string(user.ID),
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
	}, nil
}

// メールアドレスでユーザを取得
func (i *userInteractor) GetUserByEmail(ctx context.Context, email string) (*UserOutput, error) {
	user, err := i.repo.FindByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return &UserOutput{
		ID:    string(user.ID),
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

// IDでユーザを取得
func (i *userInteractor) GetUserByID(ctx context.Context, id string) (*UserOutput, error) {
	user, err := i.repo.FindByID(ctx, userdomain.UserID(id))

	if err != nil {
		return nil, err
	}

	return &UserOutput{
		ID:    string(user.ID),
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}
