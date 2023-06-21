package dto

type CreateProfileRequest struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
}

type UpdateProfileRequest struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
}
