package auth

type AuthLoginRequest struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserDTO struct {
	ID        uint    `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Phone     *string `json:"phone"`
	Email     *string `json:"email"`
	Role      string  `json:"role"`
}

type AuthLoginResponse struct {
	User  UserDTO  `json:"user"`
	Token TokenDTO `json:"token"`
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
	FirstName string  `json:"first_name" binding:"required"`
	LastName  string  `json:"last_name"`
	Email     *string `json:"email"`
	Phone     *string `json:"phone" binding:"required"`
	Password  string  `json:"password" binding:"required,min=8"`
}

type TokenDTO struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type AuthRegisterResponse struct {
	User    UserDTO `json:"user"`
	Message string  `json:"message"`
}

type AuthConfirmRequest struct {
	Phone string `json:"phone" binding:"required"`
	Otp   string `json:"otp" binding:"required"`
}

type GoogleAuthRequest struct {
	IDToken string `json:"id_token" binding:"required"`
}

func ToRegisterResponse(user *User, msg string) *AuthRegisterResponse {
	return &AuthRegisterResponse{
		User: UserDTO{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Phone:     user.Phone,
		},
		Message: msg,
	}
}
