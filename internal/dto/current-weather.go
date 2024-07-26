package dto

type CurrentWeatherDto struct {
	ApiToken string `json:"apiToken" validate:"required" description:"User API token"`
	City     string `json:"city" validate:"required" description:"City name"`
	Language string `json:"language,omitempty" description:"Language preference (optional)"`
}
