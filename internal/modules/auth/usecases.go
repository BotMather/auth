package auth

type AuthUsecase interface {
	Login() error
	Register() (*User, error)
}

type AuthUsecaseImpl struct {
	repo AuthRepository
}

func NewAuthUsecase(repo AuthRepository) AuthUsecase {
	return &AuthUsecaseImpl{
		repo: repo,
	}
}

func (a *AuthUsecaseImpl) Login() error {
	return nil
}

func (a *AuthUsecaseImpl) Register() (*User, error) {
	return nil, nil
}
