package user

type User struct {
	id           UserID
	name         UserName
	email        UserEmail
	passwordHash UserHashedPassword
	role         UserRole
}

// NewUser は、プリミティブな値から User エンティティを生成します。
// 内部で各 Value Object のバリデーションが実行されます。
func NewUser(id string, name string, email string, passwordHash string, role string) (*User, error) {
	voID, err := NewUserID(id)
	if err != nil {
		return nil, err
	}

	voName, err := NewUserName(name)
	if err != nil {
		return nil, err
	}

	voEmail, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	voPasswordHash, err := NewUserHashedPassword(passwordHash)
	if err != nil {
		return nil, err
	}

	voRole, err := NewUserRole(role)
	if err != nil {
		return nil, err
	}

	return &User{
		id:           voID,
		name:         voName,
		email:        voEmail,
		passwordHash: voPasswordHash,
		role:         voRole,
	}, nil
}

// 外部から値を取得するための Getter メソッド群
// Value Object を返すことで型安全性を担保します。
func (u *User) ID() UserID                       { return u.id }
func (u *User) Name() UserName                   { return u.name }
func (u *User) Email() UserEmail                 { return u.email }
func (u *User) Role() UserRole                   { return u.role }
func (u *User) PasswordHash() UserHashedPassword { return u.passwordHash }
