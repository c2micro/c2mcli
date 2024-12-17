package utils

import (
	"fmt"
	"time"

	"github.com/docker/go-units"
)

func HumanDuration(t time.Time) string {
	if t.IsZero() {
		return "never"
	} else {
		return units.HumanDuration(time.Since(t))
	}
}

func HumanDurationC(t time.Time) string {
	if t.IsZero() {
		return "never"
	}

	d := time.Since(t)

	days := func(hours float64) int {
		return int(hours / 24)
	}
	hours := func(value float64) int {
		return int(value)
	}
	minutes := func(minutes float64) int {
		return int(minutes)
	}
	seconds := func(seconds float64) int {
		return int(seconds)
	}

	if d.Hours() > 99*24 {
		// 100d / 222d
		return fmt.Sprintf("%dd", days(d.Hours()))
	}
	if d.Hours() > 24 {
		// 1d20h / 10d5h
		return fmt.Sprintf("%dd%dh", days(d.Hours()), hours(d.Hours())%24)
	}
	if d.Hours() > 1 {
		// 5h29m / 12h53m
		return fmt.Sprintf("%dh%dm", hours(d.Hours()), minutes(d.Minutes())%60)
	}
	if d.Minutes() > 1 {
		// 3m28s / 10m1s
		return fmt.Sprintf("%dm%ds", minutes(d.Minutes()), seconds(d.Seconds())%60)
	}
	return fmt.Sprintf("%ds", seconds(d.Seconds()))
}
