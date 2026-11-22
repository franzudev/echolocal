package config

// OpenMeteo holds the embedded weather provider, which needs no API key.
type OpenMeteo struct {
	Enabled   bool   `json:"enabled"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

const (
	DefaultOpenMeteoEnabled = true
)

func defaultOpenMeteo() OpenMeteo {
	return OpenMeteo{
		Enabled: DefaultOpenMeteoEnabled,
	}
}
