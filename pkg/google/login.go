package google

import (
	"context"
	"encoding/json"
	"time"
	"admin/pkg/request"
)

const (
	url            = "https://www.googleapis.com/oauth2/v1/userinfo?access_token="
	defaultTimeout = 10 * time.Second
)

type GoogleUser struct {
	Email         string `json:"email"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
	ID            string `json:"id"`
}

// GoogleService provides Google-related functionality
type GoogleService interface {
	GetGoogleUserInfo(ctx context.Context, token string) (*GoogleUser, error)
}

type googleService struct {
	httpClient *request.HttpClient
	secret     string
}

func NewGoogleService(httpClient *request.HttpClient, secret string) GoogleService {
	return &googleService{
		httpClient: httpClient,
		secret:     secret,
	}
}

func (s *googleService) GetGoogleUserInfo(ctx context.Context, token string) (*GoogleUser, error) {

	reqConfig := &request.Request{
		Method:  "GET",
		URL:     url + token,
		Context: ctx,
	}

	// Use the injected HttpClient to make the request
	body, err := s.httpClient.Do(reqConfig)
	if err != nil {
		return nil, err
	}

	var result GoogleUser
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
