//go:build !board_biscuit

package openmeteo

import "testing"

func TestWMO(t *testing.T) {
	tests := []struct {
		code int
		want string
	}{
		{0, "sunny"},
		{1, "partlycloudy"},
		{2, "partlycloudy"},
		{3, "cloudy"},
		{45, "fog"},
		{48, "fog"},
		{51, "rainy"},
		{53, "rainy"},
		{55, "rainy"},
		{56, "snowy-rainy"},
		{57, "snowy-rainy"},
		{61, "rainy"},
		{63, "rainy"},
		{64, "rainy"},
		{65, "pouring"},
		{66, "snowy-rainy"},
		{67, "snowy-rainy"},
		{71, "snowy"},
		{73, "snowy"},
		{75, "snowy"},
		{77, "snowy"},
		{80, "rainy"},
		{81, "rainy"},
		{82, "pouring"},
		{85, "snowy"},
		{86, "snowy"},
		{95, "lightning"},
		{96, "lightning-rainy"},
		{99, "lightning-rainy"},
		{10, "cloudy"},
		{99, "lightning-rainy"},
		{100, "cloudy"},
	}
	for _, test := range tests {
		if got := WMO(test.code); got != test.want {
			t.Errorf("WMO(%d) = %q, want %q", test.code, got, test.want)
		}
	}
}
