package dto

type RegisterRequest struct {
	FirstName  string `json:"first_name"  binding:"required"`
	SecondName string `json:"second_name" binding:"required"`
	Birthdate  string `json:"birthdate"   binding:"required"`
	Biography  string `json:"biography"`
	City       string `json:"city"        binding:"required"`
	Password   string `json:"password"    binding:"required,min=8"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
