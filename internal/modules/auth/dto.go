package auth

type AuthLoginRequest struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthLoginResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
}

type AuthMeResponse struct {
	User map[string]any `json:"user"`
}

type AuthRefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type AuthRefreshTokenResponse struct {
	AccessToken string `json:"access"`
}

type AuthRegisterRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone" binding:"required"`
	Password  string `json:"password" binding:"required,min=8"`
}
type UserDTO struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
}

type TokenDTO struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type AuthRegisterResponse struct {
	User  UserDTO  `json:"user"`
	Token TokenDTO `json:"token"`
}

func ToRegisterResponse(user *User, access string, refresh string) *AuthRegisterResponse {
	return &AuthRegisterResponse{
		User: UserDTO{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
		},
		Token: TokenDTO{
			Access:  access,
			Refresh: refresh,
		},
	}
}
