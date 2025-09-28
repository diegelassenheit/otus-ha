package dto

type RegisterUserRequest struct {
	FirstName  string `json:"first_name" binding:"required"`
	SecondName string `json:"second_name" binding:"required"`
	Birthdate  string `json:"birthdate" binding:"required"`
	Biography  string `json:"biography"`
	City       string `json:"city" binding:"required"`
	Password   string `json:"password" binding:"required,min=8"`
}

type RegisterUserResponse struct {
	UserID string `json:"user_id"`
}

type UserProfileResponse struct {
    ID         string `json:"id"`
    FirstName  string `json:"first_name"`
    SecondName string `json:"second_name"`
    Birthdate  string `json:"birthdate"`
    Biography  string `json:"biography"`
    City       string `json:"city"`
}
