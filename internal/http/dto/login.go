package dto

type LoginRequest struct {
	ID       string `json:"id"       binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
