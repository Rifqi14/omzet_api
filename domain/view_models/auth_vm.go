package view_models

type LoginVm struct {
	Token                  string `json:"token"`
	TokenExpiration        int64  `json:"token_expiration"`
	RefreshToken           string `json:"refresh_token"`
	RefreshTokenExpiration int64  `json:"refresh_token_expiration"`
}
