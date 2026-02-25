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
	hashFunc func(p string) (string, error)
}

func (m *mockHasher) Hash(p string) (string, error) { return m.hashFunc(p) }
func (m *mockHasher) Compare(h, p string) error     { return nil } // 今回は不使用

type mockIDGen struct {
	generateFunc func() userdomain.UserID
}

func (m *mockIDGen) Generate() userdomain.UserID { return m.generateFunc() }

// --- SignUp のテスト ---

func TestSignUp(t *testing.T) {
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
		"異常系: パスワードが短すぎる": {
			input:       usecase.SignUpInput{Name: "たろう", Email: "test@example.com", Password: "short"},
			prepareMock: func(r *mockUserRepo, h *mockHasher, i *mockIDGen) {},
			wantErr:     userdomain.ErrWeakPassword,
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
			// モックの初期化
			r, h, i := &mockUserRepo{}, &mockHasher{}, &mockIDGen{}
			tt.prepareMock(r, h, i)

			// 実行対象の生成（インターフェースではなく構造体を受け取る現在の実装に合わせる）
			uc := usecase.NewUserUseCase(r, h, i)

			// 実行
			out, err := uc.SignUp(ctx, tt.input)

			// 検証
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got error %v, want %v", err, tt.wantErr)
			}
			if tt.wantErr == nil {
				if out == nil {
					t.Fatal("expected output, but got nil")
				}
				if out.ID != tt.expectID {
					t.Errorf("got ID %s, want %s", out.ID, tt.expectID)
				}
			}
		})
	}
}
