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
	// NewUser コンストラクタ内で Value Object への変換とバリデーションが行われます
	newUser, err := userdomain.NewUser(
		i.idGen.Generate().String(),
		in.Name,
		in.Email,
		in.Password,
		userdomain.RoleMember.String(), // デフォルトでMemberロールを設定
	)
	if err != nil {
		return nil, err
	}

	// ユーザーを保存
	if err := i.repo.Save(ctx, newUser); err != nil {
		return nil, err
	}

	return &SignUpOutput{
		ID:    newUser.ID().String(),
		Name:  newUser.Name().String(),
		Email: newUser.Email().String(),
	}, nil
}

// ユーザのサインイン
func (i *userInteractor) SignIn(ctx context.Context, in SignInInput) (*SignInOutput, error) {
	// emailのバリデーションと正規化
	email, err := userdomain.NewEmail(in.Email)
	if err != nil {
		return nil, err
	}
	// メールアドレスでユーザを取得
	user, err := i.repo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, userdomain.ErrUserNotFound) {
			return nil, userdomain.ErrInvalidCredentials
		}
		return nil, err
	}

	// パスワードを検証
	if err := i.passwordHasher.Compare(user.PasswordHash(), in.Password); err != nil {
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
			ID:    user.ID().String(),
			Name:  user.Name().String(),
			Email: user.Email().String(),
			Role:  user.Role(),
		},
	}, nil
}

// メールアドレスでユーザを取得
func (i *userInteractor) GetUserByEmail(ctx context.Context, emails string) (*UserOutput, error) {
	// emailのバリデーションと正規化
	email, err := userdomain.NewEmail(emails)
	if err != nil {
		return nil, err
	}
	user, err := i.repo.FindByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return &UserOutput{
		ID:    user.ID().String(),
		Name:  user.Name().String(),
		Email: user.Email().String(),
		Role:  user.Role(),
	}, nil
}

// IDでユーザを取得
func (i *userInteractor) GetUserByID(ctx context.Context, id string) (*UserOutput, error) {
	userId, err := userdomain.NewUserID(id)
	if err != nil {
		return nil, err
	}
	user, err := i.repo.FindByID(ctx, userId)

	if err != nil {
		return nil, err
	}

	return &UserOutput{
		ID:    user.ID().String(),
		Name:  user.Name().String(),
		Email: user.Email().String(),
		Role:  user.Role(),
	}, nil
}
