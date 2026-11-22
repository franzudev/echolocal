//go:build !board_biscuit

package openmeteo

// WMO turns an Open-Meteo weather code into the condition word the shared
// dashboard weather view draws (the Home Assistant condition vocabulary).
func WMO(n int) string {
	switch {
	case n == 0:
		return "sunny"
	case n == 1 || n == 2:
		return "partlycloudy"
	case n == 3:
		return "cloudy"
	case n == 45 || n == 48:
		return "fog"
	case n >= 51 && n <= 55: // drizzle
		return "rainy"
	case n == 56 || n == 57: // freezing drizzle
		return "snowy-rainy"
	case n >= 61 && n <= 64: // rain
		return "rainy"
	case n == 65: // heavy rain
		return "pouring"
	case n == 66 || n == 67: // freezing rain
		return "snowy-rainy"
	case n >= 71 && n <= 77: // snow
		return "snowy"
	case n == 80 || n == 81: // showers
		return "rainy"
	case n == 82: // violent showers
		return "pouring"
	case n == 85 || n == 86: // snow showers
		return "snowy"
	case n == 95: // thunderstorm
		return "lightning"
	case n == 96 || n == 99: // thunderstorm with hail
		return "lightning-rainy"
	}
	return "cloudy"
}
