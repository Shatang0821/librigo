package usecase_test

import (
	"context"
	"errors"
	userdomain "librigo/internal/domain/user"
	"librigo/internal/usecase"
	"testing"
)

// --- 手動モックの定義 ---

type mockUserRepo struct {
	saveFunc        func(u *userdomain.User) error
	findByEmailFunc func(email string) (*userdomain.User, error)
	findByIDFunc    func(id userdomain.UserID) (*userdomain.User, error)
}

func (m *mockUserRepo) Save(ctx context.Context, u *userdomain.User) error { return m.saveFunc(u) }
func (m *mockUserRepo) FindByEmail(ctx context.Context, email string) (*userdomain.User, error) {
	return m.findByEmailFunc(email)
}
func (m *mockUserRepo) FindByID(ctx context.Context, id userdomain.UserID) (*userdomain.User, error) {
	return m.findByIDFunc(id)
}

type mockHasher struct {
	hashFunc    func(p string) (string, error)
	compareFunc func(h, p string) error
}

func (m *mockHasher) Hash(p string) (string, error) { return m.hashFunc(p) }
func (m *mockHasher) Compare(h, p string) error     { return m.compareFunc(h, p) }

type mockIDGen struct {
	generateFunc func() userdomain.UserID
}

func (m *mockIDGen) Generate() userdomain.UserID { return m.generateFunc() }

// --- SignUp のテスト ---

func TestUserUseCase_SignUp(t *testing.T) {
	ctx := context.Background()
	fixedID := "uuid-123"

	tests := map[string]struct {
		input       usecase.SignUpInput
		prepareMock func(r *mockUserRepo, h *mockHasher, i *mockIDGen)
		wantErr     error
		expectID    string
	}{
		"正常系: ユーザー登録成功": {
			input: usecase.SignUpInput{Name: "たろう", Email: "test@example.com", Password: "password123"},
			prepareMock: func(r *mockUserRepo, h *mockHasher, i *mockIDGen) {
				r.findByEmailFunc = func(email string) (*userdomain.User, error) { return nil, nil }
				r.saveFunc = func(u *userdomain.User) error { return nil }
				h.hashFunc = func(p string) (string, error) { return "hashed_pass", nil }
				i.generateFunc = func() userdomain.UserID { return userdomain.UserID(fixedID) }
			},
			wantErr:  nil,
			expectID: fixedID,
		},
		"異常系: Emailが既に存在する": {
			input: usecase.SignUpInput{Name: "たろう", Email: "dup@example.com", Password: "password123"},
			prepareMock: func(r *mockUserRepo, h *mockHasher, i *mockIDGen) {
				r.findByEmailFunc = func(email string) (*userdomain.User, error) {
					return &userdomain.User{Email: email}, nil
				}
			},
			wantErr: userdomain.ErrDuplicateUser,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			r, h, i := &mockUserRepo{}, &mockHasher{}, &mockIDGen{}
			tt.prepareMock(r, h, i)
			uc := usecase.NewUserUseCase(r, h, i)

			out, err := uc.SignUp(ctx, tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got error %v, want %v", err, tt.wantErr)
			}
			if tt.wantErr == nil && out.ID != tt.expectID {
				t.Errorf("got ID %s, want %s", out.ID, tt.expectID)
			}
		})
	}
}

func TestUserUseCase_SignIn(t *testing.T) {
	ctx := context.Background()

	tests := map[string]struct {
		input       usecase.SignInInput
		prepareMock func(r *mockUserRepo, h *mockHasher)
		wantErr     error
	}{
		"正常系: ログイン成功": {
			input: usecase.SignInInput{Email: "test@example.com", Password: "password123"},
			prepareMock: func(r *mockUserRepo, h *mockHasher) {
				r.findByEmailFunc = func(email string) (*userdomain.User, error) {
					return &userdomain.User{Email: email, PasswordHash: "hashed"}, nil
				}
				h.compareFunc = func(h, p string) error { return nil }
			},
			wantErr: nil,
		},
		"異常系: ユーザーが見つからない": {
			input: usecase.SignInInput{Email: "none@example.com", Password: "password123"},
			prepareMock: func(r *mockUserRepo, h *mockHasher) {
				r.findByEmailFunc = func(email string) (*userdomain.User, error) {
					return nil, userdomain.ErrUserNotFound
				}
			},
			wantErr: userdomain.ErrInvalidCredentials,
		},
		"異常系: パスワード不一致": {
			input: usecase.SignInInput{Email: "test@example.com", Password: "wrong"},
			prepareMock: func(r *mockUserRepo, h *mockHasher) {
				r.findByEmailFunc = func(email string) (*userdomain.User, error) {
					return &userdomain.User{Email: email}, nil
				}
				h.compareFunc = func(h, p string) error { return errors.New("mismatch") }
			},
			wantErr: userdomain.ErrInvalidCredentials,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			r, h, i := &mockUserRepo{}, &mockHasher{}, &mockIDGen{}
			tt.prepareMock(r, h)
			uc := usecase.NewUserUseCase(r, h, i)

			out, err := uc.SignIn(ctx, tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got error %v, want %v", err, tt.wantErr)
			}
			if tt.wantErr == nil && out.Token == "" {
				t.Error("expected token, but got empty")
			}
		})
	}
}

// --- 取得系のテスト ---

func TestUserUseCase_GetMethods(t *testing.T) {
	ctx := context.Background()
	testUser := &userdomain.User{ID: "u1", Name: "たろう", Email: "test@example.com"}

	r, h, i := &mockUserRepo{}, &mockHasher{}, &mockIDGen{}
	uc := usecase.NewUserUseCase(r, h, i)

	t.Run("GetUserByEmail: 成功", func(t *testing.T) {
		r.findByEmailFunc = func(email string) (*userdomain.User, error) { return testUser, nil }
		out, err := uc.GetUserByEmail(ctx, testUser.Email)
		if err != nil || out.ID != "u1" {
			t.Errorf("failed to get user: %v", err)
		}
	})

	t.Run("GetUserByID: 成功", func(t *testing.T) {
		r.findByIDFunc = func(id userdomain.UserID) (*userdomain.User, error) { return testUser, nil }
		out, err := uc.GetUserByID(ctx, "u1")
		if err != nil || out.Email != testUser.Email {
			t.Errorf("failed to get user: %v", err)
		}
	})
}
