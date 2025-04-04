package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	loginURL = "http://localhost:9999/auth/login"
	hiURL    = "http://localhost:9999/hi"
)

type responseLogin struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type apiClient struct {
	client *http.Client
}

func newAPIClient(timeout time.Duration) *apiClient {
	return &apiClient{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (a *apiClient) login(username, password string) (*responseLogin, error) {
	loginPayload := map[string]string{
		"username": username,
		"password": password,
	}
	
	payloadBytes, err := json.Marshal(loginPayload)
	if err != nil {
		return nil, fmt.Errorf("error marshaling login payload: %w", err)
	}

	req, err := http.NewRequest("POST", loginURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("error creating login request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Device-Id", "APP")

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending login request: %w", err)
	}
	defer resp.Body.Close()

	var responseData responseLogin
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return nil, fmt.Errorf("error decoding login response: %w", err)
	}

	return &responseData, nil
}

func (a *apiClient) getHi(token string) (string, error) {
	req, err := http.NewRequest("GET", hiURL, nil)
	if err != nil {
		return "", fmt.Errorf("error creating hi request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := a.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending hi request: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading hi response: %w", err)
	}

	return string(bodyBytes), nil
}

func main() {
	client := newAPIClient(10 * time.Second)
	
	// Login
	auth, err := client.login("admin", "admin")
	if err != nil {
		fmt.Println("Login error:", err)
		return
	}

	fmt.Println("Authentication successful")
	fmt.Println("Access Token:", auth.AccessToken)
	fmt.Println("Refresh Token:", auth.RefreshToken)

	// Call Hi API
	hiResponse, err := client.getHi(auth.AccessToken)
	if err != nil {
		fmt.Println("Hi API error:", err)
		return
	}

	fmt.Println("Result get hi:", hiResponse)
}
