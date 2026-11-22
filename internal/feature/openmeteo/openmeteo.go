//go:build !board_biscuit

package openmeteo

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/franzudev/libcountertop/pkg/fetch"
	haweather "github.com/franzudev/libcountertop/pkg/homeassistant/weather"
	sharedcomponent "github.com/franzudev/libcountertop/pkg/runtime/component"
	"github.com/franzudev/libcountertop/pkg/runtime/service"
)

const (
	endpoint = "https://api.open-meteo.com/v1/forecast"
	every    = 15 * time.Minute
	wait     = 20 * time.Second
)

type Forecast struct {
	Current struct {
		Temperature float64 `json:"temperature_2m"`
		Apparent    float64 `json:"apparent_temperature"`
		Code        int     `json:"weather_code"`
		Wind        float64 `json:"wind_speed_10m"`
		Humidity    float64 `json:"relative_humidity_2m"`
	} `json:"current"`
}

type Options struct {
	Enabled func() bool
	LatText func() string
	LonText func() string
	Client  *http.Client
	Agent   string
	Redraw  func()
}

type Source struct {
	o Options

	mu       sync.Mutex
	reading  haweather.Reading
	ok       bool
	at       time.Time
	fetching bool
}

func New(o Options) *Source {
	if o.Client == nil {
		o.Client = fetch.Client(wait)
	}
	return &Source{o: o}
}

func (s *Source) Name() string { return "openmeteo" }

func (s *Source) Run(ctx context.Context) error {
	t := time.NewTimer(30 * time.Second)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-t.C:
		}
		if at := s.coords(); s.enabled() && at != nil {
			s.poll(ctx, at[0], at[1])
		}
		t.Reset(every)
	}
}

func (s *Source) enabled() bool {
	return s.o.Enabled == nil || s.o.Enabled()
}

// coords returns the configured location, or nil when unset.
func (s *Source) coords() *[2]float64 {
	if s.o.LatText == nil || s.o.LonText == nil {
		return nil
	}
	lat, lon := s.o.LatText(), s.o.LonText()
	if lat == "" || lon == "" {
		return nil
	}
	l, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return nil
	}
	n, err := strconv.ParseFloat(lon, 64)
	if err != nil {
		return nil
	}
	if math.Abs(l) > 90 || math.Abs(n) > 180 {
		return nil
	}
	return &[2]float64{l, n}
}

func (s *Source) poll(ctx context.Context, lat, lon float64) {
	ctx, cancel := context.WithTimeout(ctx, wait)
	defer cancel()
	s.mu.Lock()
	if s.fetching {
		s.mu.Unlock()
		return
	}
	s.fetching = true
	s.mu.Unlock()

	reading, err := s.fetch(ctx, lat, lon)
	s.mu.Lock()
	s.fetching = false
	if err == nil {
		s.reading, s.ok, s.at = reading, true, time.Now()
	}
	s.mu.Unlock()
	if err != nil {
		slog.Warn("reading the open-meteo forecast", "err", err)
		return
	}
	s.redraw()
}

func (s *Source) fetch(ctx context.Context, lat, lon float64) (haweather.Reading, error) {
	url := fmt.Sprintf(
		"%s?latitude=%.4f&longitude=%.4f&current=temperature_2m,apparent_temperature,weather_code,wind_speed_10m,relative_humidity_2m&forecast_days=5&timezone=auto",
		endpoint, lat, lon,
	)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return haweather.Reading{}, fmt.Errorf("open-meteo: %w", err)
	}
	req.Header.Set("User-Agent", s.o.Agent)
	req.Header.Set("Accept", "application/json")
	resp, err := s.o.Client.Do(req)
	if err != nil {
		return haweather.Reading{}, fmt.Errorf("open-meteo: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return haweather.Reading{}, fmt.Errorf("open-meteo: %s", resp.Status)
	}
	var f Forecast
	if err := json.NewDecoder(http.MaxBytesReader(nil, resp.Body, 4<<20)).Decode(&f); err != nil {
		return haweather.Reading{}, fmt.Errorf("open-meteo: the forecast: %w", err)
	}

	temperature, apparent := f.Current.Temperature, f.Current.Apparent
	wind, humidity := f.Current.Wind, f.Current.Humidity
	return haweather.Reading{
		Condition:   WMO(f.Current.Code),
		Temperature: &temperature,
		Unit:        "°C",
		Feels:       &apparent,
		Wind:        &wind,
		WindUnit:    "km/h",
		Humidity:    &humidity,
	}, nil
}

func (s *Source) redraw() {
	if s.o.Redraw != nil {
		s.o.Redraw()
	}
}

// fresh says the reading is present and recent enough to trust.
func (s *Source) fresh() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.ok && time.Since(s.at) < 2*every
}

// Now reports a reading only when Open-Meteo is enabled, located and fresh.
func (s *Source) Now() (haweather.Reading, bool) {
	if !s.enabled() || !s.fresh() {
		return haweather.Reading{}, false
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.reading, true
}

var _ sharedcomponent.Runner = (*Source)(nil)
var _ service.Service = (*Source)(nil)
