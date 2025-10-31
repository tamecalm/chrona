package tz

import (
	"fmt"

	"github.com/ipinfo/go/v2/ipinfo"
)

func DetectLocation() (country, state string, err error) {
	client := ipinfo.NewClient(nil, nil, "")
	info, err := client.GetIPInfo(nil)
	if err != nil {
		return "", "", err
	}
	country = info.CountryName
	state = info.Region
	fmt.Printf("🌐 Detected: %s, %s\n", state, country)
	return
}
