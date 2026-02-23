package user

type PassWordHasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}
