package tz

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"chrona/internal/ui"
)

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func ShowTime(country, state string) {
	states := Locations[country]
	var location *Location
	for _, s := range states {
		if strings.EqualFold(s.State, state) {
			location = &s
			break
		}
	}
	if location == nil {
		fmt.Println(ui.Error("❌ Invalid state"))
		return
	}

	loc, err := time.LoadLocation(location.Zone)
	if err != nil {
		fmt.Println(ui.Error("❌ Failed to load timezone"))
		return
	}

	weather, weatherErr := FetchWeather(location.Lat, location.Lon)

	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

	quit := make(chan bool)
	keyPress := make(chan byte)

	go func() {
		var b []byte = make([]byte, 1)
		for {
			os.Stdin.Read(b)
			keyPress <- b[0]
		}
	}()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	lastSecond := -1

	for {
		select {
		case <-ticker.C:
			now := time.Now().In(loc)

			if now.Second() == lastSecond {
				continue
			}
			lastSecond = now.Second()

			clearScreen()

			fmt.Println(ui.Title("╔═══════════════════════════════════════════════════════════════════╗"))
			fmt.Println(ui.Title("║              🌍  CHRONA - REAL-TIME TIMEZONE CLOCK  🌍            ║"))
			fmt.Println(ui.Title("╚═══════════════════════════════════════════════════════════════════╝"))
			fmt.Println()

			fmt.Printf("  %s %s\n", ui.Label("Country:"), ui.Value(strings.Title(country)))
			fmt.Printf("  %s %s\n", ui.Label("State/City:"), ui.Value(state))
			fmt.Printf("  %s %s\n", ui.Label("Timezone:"), ui.Value(location.Zone))
			fmt.Println()

			fmt.Println(ui.Title("┌─────────────────────────────────────────────────────────────────┐"))
			fmt.Printf("│  ⏰  %s                                           │\n", ui.Value(now.Format("03:04:05 PM")))
			fmt.Printf("│  📅  %s                              │\n", ui.Value(now.Format("Monday, January 02, 2006")))
			fmt.Println(ui.Title("└─────────────────────────────────────────────────────────────────┘"))
			fmt.Println()

			if weatherErr == nil && weather != nil {
				fmt.Println(ui.Title("┌─────────────────────────────────────────────────────────────────┐"))
				fmt.Println(ui.Title("│                        WEATHER INFORMATION                       │"))
				fmt.Println(ui.Title("├─────────────────────────────────────────────────────────────────┤"))

				dayNight := "🌙 Night"
				if weather.IsDay {
					dayNight = "☀️  Day"
				}

				fmt.Printf("│  %s %s                                            │\n", ui.Label("Condition:"), ui.Value(weather.Condition))
				fmt.Printf("│  %s %s°C / %.1f°F                                  │\n",
					ui.Label("Temperature:"),
					ui.Value(fmt.Sprintf("%.1f", weather.Temperature)),
					weather.Temperature*9/5+32)
				fmt.Printf("│  %s %s                                          │\n", ui.Label("Period:"), ui.Value(dayNight))

				sunriseTime := ""
				sunsetTime := ""
				if len(weather.Sunrise) >= 16 {
					if t, err := time.Parse(time.RFC3339, weather.Sunrise); err == nil {
						sunriseTime = t.In(loc).Format("03:04 PM")
					}
				}
				if len(weather.Sunset) >= 16 {
					if t, err := time.Parse(time.RFC3339, weather.Sunset); err == nil {
						sunsetTime = t.In(loc).Format("03:04 PM")
					}
				}

				if sunriseTime != "" && sunsetTime != "" {
					fmt.Printf("│  %s %s      %s %s                │\n",
						ui.Label("🌅 Sunrise:"), ui.Value(sunriseTime),
						ui.Label("🌇 Sunset:"), ui.Value(sunsetTime))
				}

				fmt.Println(ui.Title("└─────────────────────────────────────────────────────────────────┘"))
			} else {
				fmt.Println(ui.Info("┌─────────────────────────────────────────────────────────────────┐"))
				fmt.Println(ui.Info("│  Weather information unavailable                                 │"))
				fmt.Println(ui.Info("└─────────────────────────────────────────────────────────────────┘"))
			}

			fmt.Println()
			fmt.Println(ui.Info("  Press [B] to go back  |  Press [Q] to quit"))
			fmt.Println()

		case key := <-keyPress:
			keyUpper := strings.ToUpper(string(key))
			if keyUpper == "B" {
				fmt.Println()
				return
			} else if keyUpper == "Q" {
				clearScreen()
				fmt.Println(ui.Success("\n  👋 Thank you for using Chrona. Goodbye!\n"))
				os.Exit(0)
			}

		case <-quit:
			return
		}
	}
}
