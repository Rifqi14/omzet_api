package request

type LoginRequest struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}
