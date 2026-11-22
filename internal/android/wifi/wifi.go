package wifi

import (
	"github.com/franzudev/libcountertop/pkg/host/adb"
	"github.com/franzudev/libcountertop/pkg/host/wifisetup"

	"github.com/franzudev/echolocal/internal/layout"
)

const dhcpService = "dhcpcd-wlan0"

func On(d *adb.Device) wifisetup.Supplicant {
	return wifisetup.Supplicant{
		Device:  d,
		Sockets: layout.WifiSockets,
		Iface:   layout.WifiIface,
		DHCP:    dhcpService,
	}
}
