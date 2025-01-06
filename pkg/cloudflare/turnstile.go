package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
	"admin/pkg/request"
)

const (
	turnstileVerifyURL = "https://challenges.cloudflare.com/turnstile/v0/siteverify"
	defaultTimeout     = 10 * time.Second
)

// TurnstileResponse represents the response from Cloudflare Turnstile verification
type TurnstileResponse struct {
	Success    bool     `json:"success"`
	ErrorCodes []string `json:"error-codes,omitempty"`
	Challenge  string   `json:"challenge_ts,omitempty"`
	Hostname   string   `json:"hostname,omitempty"`
	Action     string   `json:"action,omitempty"`
	CData      string   `json:"cdata,omitempty"`
}

// Client represents a Cloudflare Turnstile verification client
type Client struct {
	httpClient *request.HttpClient
	secret     string
}

// NewClient creates a new Cloudflare Turnstile client
func NewClient(client *request.HttpClient, secret string) *Client {
	// Add default timeout option
	return &Client{
		httpClient: client,
		secret:     secret,
	}
}

// VerifyRequest represents the parameters for verification
type VerifyRequest struct {
	// Token from the client-side widget
	Token string
	// Optional IP address of the user
	RemoteIP string
}

// Verify performs Turnstile token verification
func (c *Client) Verify(ctx context.Context, req *VerifyRequest) (*TurnstileResponse, error) {
	if req.Token == "" {
		return nil, fmt.Errorf("token cannot be empty")
	}

	// Prepare form data
	formData := url.Values{}
	formData.Set("secret", c.secret)
	formData.Set("response", req.Token)

	if req.RemoteIP != "" {
		formData.Set("remoteip", req.RemoteIP)
	}

	// Prepare headers
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	// Create request config
	reqConfig := &request.Request{
		Method:  "POST",
		URL:     turnstileVerifyURL,
		Headers: headers,
		Body:    formData,
		Context: ctx,
	}

	// Execute request
	body, err := c.httpClient.Do(reqConfig)
	if err != nil {
		return nil, fmt.Errorf("turnstile verification request failed: %w", err)
	}

	// Parse response
	var response TurnstileResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse turnstile response: %w", err)
	}

	if !response.Success {
		return &response, fmt.Errorf("turnstile verification failed: %v", response.ErrorCodes)
	}

	return &response, nil
}
