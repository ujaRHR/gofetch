package info

import (
	"fmt"
	"os"
	"strings"
)

func mapDesktopToWM(d string) string {
	s := strings.ToLower(d)

	switch {
	case strings.Contains(s, "kde"):
		return "KWin"
	case strings.Contains(s, "gnome"):
		return "Mutter"
	case strings.Contains(s, "xfce"):
		return "Xfwm4"
	case strings.Contains(s, "cinnamon"):
		return "Muffin"
	case strings.Contains(s, "mate"):
		return "Marco"
	case strings.Contains(s, "lxqt"), strings.Contains(s, "lxde"):
		return "Openbox"
	case strings.Contains(s, "sway"):
		return "Sway"
	case strings.Contains(s, "hyprland"):
		return "Hyprland"

	default:
		return "Unknown"
	}
}

func GetWMInfo() string {
	protocol := func() string {
		if os.Getenv("WAYLAND_DISPLAY") != "" {
			return "wayland"
		}
		if os.Getenv("DISPLAY") != "" {
			return "x11"
		}
		if v := os.Getenv("XDG_SESSION_TYPE"); v != "" {
			return strings.ToTitle(v)
		}
		return "Unknown"
	}()

	var wm string

	if v := os.Getenv("XDG_CURRENT_DESKTOP"); v != "" {
		wm = mapDesktopToWM(v)
	}

	return fmt.Sprintf("%s (%s)", wm, protocol)
}
