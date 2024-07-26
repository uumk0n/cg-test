package services

import (
	"cg-test/config"
	"cg-test/github.com/uumkon/cg-test/proto/weather"
	"cg-test/internal/entities"
	"cg-test/internal/storage/repo"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type WeatherServiceServer struct {
	userRepo    *repo.UserRepo
	actionsRepo *repo.ActionsRepo
	config      *config.Config
	weather.UnimplementedWeatherServiceServer
}

func NewWeatherServiceServer(userRepo *repo.UserRepo, actionsRepo *repo.ActionsRepo, config *config.Config) *WeatherServiceServer {
	return &WeatherServiceServer{userRepo: userRepo, actionsRepo: actionsRepo, config: config}
}

func (s *WeatherServiceServer) GetCurrentWeather(ctx context.Context, req *weather.CurrentWeatherRequest) (*weather.CurrentWeatherResponse, error) {
	apiToken, city, language := req.GetApiToken(), req.GetCity(), req.GetLanguage()

	user, err := s.userRepo.FindUserByApiToken(apiToken)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	if user.APIToken == "" {
		return nil, errors.New("invalid token")
	}

	weatherResponse, err := s.fetchWeatherData(city, language)
	if err != nil {
		s.saveAction(user.ID, time.Now(), 0, nil)
		return nil, err
	}
	s.saveAction(user.ID, time.Now(), weatherResponse.StatusCode, weatherResponse.TempC)

	var tempC float32
	if weatherResponse.TempC != nil {
		tempC = float32(*weatherResponse.TempC)
	}

	dataJSON, err := json.Marshal(weatherResponse.Data)
	if err != nil {
		return nil, err
	}

	return &weather.CurrentWeatherResponse{
		StatusCode: int32(weatherResponse.StatusCode),
		TempC:      tempC,
		Data:       string(dataJSON),
	}, nil
}

func (s *WeatherServiceServer) fetchWeatherData(city, language string) (*WeatherResponse, error) {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&lang=%s", s.config.WeatherAPI.ApiKey, city, language)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error fetching weather data")
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	tempC := result["current"].(map[string]interface{})["temp_c"].(float64)
	return &WeatherResponse{
		StatusCode: resp.StatusCode,
		TempC:      &tempC,
		Data:       result,
	}, nil
}

func (s *WeatherServiceServer) saveAction(userID int, actionTime time.Time, requestResult int, tempC *float64) {
	action := entities.Actions{
		UserID:        userID,
		ActionTime:    actionTime,
		RequestResult: requestResult,
		TempC:         float64(*tempC),
	}
	s.actionsRepo.SaveAction(action)
}

type WeatherResponse struct {
	StatusCode int
	TempC      *float64
	Data       map[string]interface{}
}
