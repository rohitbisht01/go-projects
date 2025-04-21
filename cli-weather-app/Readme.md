# 🌤️ GoWeather CLI

A fast and lightweight CLI-based weather application written in Go. It allows users to fetch real-time weather data for any city using the OpenWeatherMap API — right from their terminal.

---

## Features

- Fetch weather data by city name
- Displays formatted output in a clean table
- Uses OpenWeatherMap API
- CLI spinner while fetching data
- `.env` support for API keys

---

## Getting Started

### 1. Clone the Repository

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Setup Environment Variables

Create a .env file in the root directory with your OpenWeatherMap API key:

```bash
WEATHER_API_KEY=your_api_key_here
```

### Usage

To fetch weather data for a city:

```bash
go run ./ -fetch Delhi
```

This will display the current weather data for Delhi in a neat table format.

### CLI Flags

| Flag     | Description                   | Example         |
| -------- | ----------------------------- | --------------- |
| `-fetch` | Fetch weather data for a city | `-fetch London` |
