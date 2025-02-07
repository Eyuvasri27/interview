package service

import (
	"context"
	"country-api/cache"
	"country-api/client"
)

type CountryService struct {
	Cache  cache.Cache
	Client *client.RestCountriesClient
}

func NewCountryService(cache cache.Cache, client *client.RestCountriesClient) *CountryService {
	return &CountryService{
		Cache:  cache,
		Client: client,
	}
}

func (s *CountryService) GetCountryByName(ctx context.Context, name string) (*client.Country, error) {
	// Check cache first
	if data, exists := s.Cache.Get(name); exists {
		return data.(*client.Country), nil
	}

	// Fetch from REST Countries API
	country, err := s.Client.GetCountryByName(ctx, name)
	if err != nil {
		return nil, err
	}

	// Cache the result
	s.Cache.Set(name, country)

	return country, nil
}
