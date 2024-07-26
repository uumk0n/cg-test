package dto

type LoginDto struct {
	Login    string `json:"login" validate:"required" description:"User login"`
	Password string `json:"password" validate:"required" description:"User password"`
}
