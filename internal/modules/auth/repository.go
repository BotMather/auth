package auth

type AuthRepository interface {
	GetID(int64) *User
}

type AuthRepositoryImpl struct{}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (a *AuthRepositoryImpl) GetID(id int64) *User {
	return &User{}
}
