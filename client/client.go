package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Country struct {
	Name       Name                `json:"name"`
	Capital    []string            `json:"capital"`
	Currencies map[string]Currency `json:"currencies"`
	Population int                 `json:"population"`
}

type Name struct {
	Common   string `json:"common"`
	Official string `json:"official"`
}

type Currency struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type RestCountriesClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewRestCountriesClient() *RestCountriesClient {
	return &RestCountriesClient{
		BaseURL: "https://restcountries.com/v3.1",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *RestCountriesClient) GetCountryByName(ctx context.Context, name string) (*Country, error) {
	url := fmt.Sprintf("%s/name/%s", c.BaseURL, name)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch country data: %s", resp.Status)
	}

	var countries []Country
	if err := json.NewDecoder(resp.Body).Decode(&countries); err != nil {
		return nil, err
	}

	if len(countries) == 0 {
		return nil, fmt.Errorf("country not found")
	}

	return &countries[0], nil
}
