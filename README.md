# interview# Country Search API

This is a REST API service built in Go that provides country information using the [REST Countries API](https://restcountries.com/). It includes a custom in-memory cache and unit tests.

## Features

- Single endpoint: `/api/countries/search`
- Query parameter: `name` (string)
- Custom in-memory cache implementation
- Graceful shutdown
- Unit tests for cache, HTTP client, service, and handlers with ECO framework

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/Eyuvasri27/interview.git

## Run the application:

   go run main.go