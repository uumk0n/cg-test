package dto

import "regexp"

type RegisterDto struct {
	Login    string `json:"login" validate:"required" description:"User login"`
	Password string `json:"password" validate:"required,min=6,regex=^[a-zA-Z0-9.,!_]+$" description:"User password (must be at least 6 characters long and contain one of the characters: . , ! _)"`
	FIO      string `json:"fio" validate:"required" description:"Full name of the user"`
}

// ValidatePassword ensures the password meets the required criteria.
func ValidatePassword(password string) bool {
	re := regexp.MustCompile(`[.,!_]`)
	return len(password) >= 6 && re.MatchString(password)
}
