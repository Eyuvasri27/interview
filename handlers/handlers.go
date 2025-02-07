package handlers

import (
	"context"
	"country-api/service"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type CountryHandler struct {
	Service *service.CountryService
}

func (h *CountryHandler) SearchCountry(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	name := c.QueryParam("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "name parameter is required"})
	}

	country, err := h.Service.GetCountryByName(ctx, name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Extract currency information
	currencyInfo := make(map[string]string)
	for code, currency := range country.Currencies {
		currencyInfo[code] = currency.Name + " (" + currency.Symbol + ")"
	}

	// Return the response
	response := map[string]interface{}{
		"name":       country.Name.Common,
		"capital":    country.Capital,
		"currency":   currencyInfo,
		"population": country.Population,
	}

	return c.JSON(http.StatusOK, response)
}
