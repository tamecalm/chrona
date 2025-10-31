package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"chrona/internal/tz"
	"chrona/internal/ui"
	"github.com/AlecAivazis/survey/v2"
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

func main() {
	for {
		clearScreen()
		fmt.Println(ui.Title("\n╔═══════════════════════════════════════════════════════════════════╗"))
		fmt.Println(ui.Title("║              🌍  CHRONA - REAL-TIME TIMEZONE CLOCK  🌍            ║"))
		fmt.Println(ui.Title("╚═══════════════════════════════════════════════════════════════════╝"))
		fmt.Println()

		var useGeo bool
		confirm := &survey.Confirm{
			Message: "Would you like to auto-detect your location?",
			Default: true,
		}
		_ = survey.AskOne(confirm, &useGeo)

		if useGeo {
			country, state, err := tz.DetectLocation()
			if err == nil && country != "" {
				fmt.Println(ui.Success(fmt.Sprintf("\n✅ Successfully detected: %s, %s\n", state, country)))
				time.Sleep(1 * time.Second)
				tz.ShowTime(strings.ToLower(country), state)
				continue
			}
			fmt.Println(ui.Error("\n❌ Auto-detection failed. Please select manually.\n"))
			time.Sleep(1 * time.Second)
		}

		var countryOptions []string
		for c := range tz.Locations {
			countryOptions = append(countryOptions, strings.Title(c))
		}
		countryOptions = append(countryOptions, "Exit")

		var country string
		countryPrompt := &survey.Select{
			Message: "Select a country:",
			Options: countryOptions,
		}
		_ = survey.AskOne(countryPrompt, &country)

		if country == "Exit" {
			clearScreen()
			fmt.Println(ui.Success("\n  👋 Thank you for using Chrona. Goodbye!\n"))
			return
		}

		country = strings.ToLower(country)
		states := tz.Locations[country]

		var state string
		if len(states) > 1 {
			var stateOptions []string
			for _, s := range states {
				stateOptions = append(stateOptions, s.State)
			}
			stateOptions = append(stateOptions, "← Back", "Exit")

			statePrompt := &survey.Select{
				Message: fmt.Sprintf("Select a state/city in %s:", strings.Title(country)),
				Options: stateOptions,
			}
			_ = survey.AskOne(statePrompt, &state)

			if state == "← Back" {
				continue
			}
			if state == "Exit" {
				clearScreen()
				fmt.Println(ui.Success("\n  👋 Thank you for using Chrona. Goodbye!\n"))
				return
			}
		} else {
			state = states[0].State
		}

		tz.ShowTime(country, state)
	}
}
