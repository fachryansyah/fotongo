package dtos

type LoginRequest struct {
	Email    string `form:"email" json:"email" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

type LoginResponse struct {
	TokenJWT string `json:"token_jwt"`
}

type RegisterRequest struct {
	Username string `form:"username" json:"username" validate:"required"`
	Email    string `form:"email" json:"email" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

type RegisterResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type CheckIfRegisteredRequest struct {
	GoogleID string `form:"google_id" json:"google_id" validate:"required"`
}

type CheckIfRegisteredResponse struct {
	Registered bool `json:"is_registered"`
}
