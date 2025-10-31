package tz

type Location struct {
	State string
	Zone  string
	Lat   float64
	Lon   float64
}

var Locations = map[string][]Location{
	"usa": {
		{"New York", "America/New_York", 40.7128, -74.0060},
		{"Florida", "America/New_York", 27.9944, -81.7603},
		{"Georgia", "America/New_York", 32.1656, -82.9001},
		{"Washington D.C.", "America/New_York", 38.9072, -77.0369},
		{"Chicago (Illinois)", "America/Chicago", 41.8781, -87.6298},
		{"Dallas (Texas)", "America/Chicago", 32.7767, -96.7970},
		{"Houston (Texas)", "America/Chicago", 29.7604, -95.3698},
		{"El Paso (Texas)", "America/Denver", 31.7619, -106.4850},
		{"Denver (Colorado)", "America/Denver", 39.7392, -104.9903},
		{"Phoenix (Arizona)", "America/Phoenix", 33.4484, -112.0740},
		{"Los Angeles (California)", "America/Los_Angeles", 34.0522, -118.2437},
		{"Seattle (Washington)", "America/Los_Angeles", 47.6062, -122.3321},
		{"Anchorage (Alaska)", "America/Anchorage", 61.2181, -149.9003},
		{"Honolulu (Hawaii)", "Pacific/Honolulu", 21.3099, -157.8581},
	},
	"nigeria": {
		{"Lagos", "Africa/Lagos", 6.5244, 3.3792},
		{"Abuja", "Africa/Lagos", 9.0765, 7.3986},
	},
	"canada": {
		{"Toronto (Ontario)", "America/Toronto", 43.6532, -79.3832},
		{"Vancouver (British Columbia)", "America/Vancouver", 49.2827, -123.1207},
		{"Edmonton (Alberta)", "America/Edmonton", 53.5461, -113.4938},
		{"Halifax (Nova Scotia)", "America/Halifax", 44.6488, -63.5752},
	},
	"india": {
		{"New Delhi", "Asia/Kolkata", 28.6139, 77.2090},
		{"Mumbai", "Asia/Kolkata", 19.0760, 72.8777},
		{"Bangalore", "Asia/Kolkata", 12.9716, 77.5946},
	},
	"china": {
		{"Beijing", "Asia/Shanghai", 39.9042, 116.4074},
		{"Shanghai", "Asia/Shanghai", 31.2304, 121.4737},
		{"Hong Kong", "Asia/Hong_Kong", 22.3193, 114.1694},
	},
	"uk": {
		{"London", "Europe/London", 51.5074, -0.1278},
		{"Manchester", "Europe/London", 53.4808, -2.2426},
	},
	"australia": {
		{"Sydney", "Australia/Sydney", -33.8688, 151.2093},
		{"Melbourne", "Australia/Melbourne", -37.8136, 144.9631},
		{"Adelaide", "Australia/Adelaide", -34.9285, 138.6007},
		{"Perth", "Australia/Perth", -31.9505, 115.8605},
	},
	"brazil": {
		{"São Paulo", "America/Sao_Paulo", -23.5505, -46.6333},
		{"Rio de Janeiro", "America/Sao_Paulo", -22.9068, -43.1729},
		{"Manaus", "America/Manaus", -3.1190, -60.0217},
		{"Salvador", "America/Bahia", -12.9714, -38.5014},
	},
	"south africa": {
		{"Johannesburg", "Africa/Johannesburg", -26.2041, 28.0473},
		{"Cape Town", "Africa/Johannesburg", -33.9249, 18.4241},
	},
}
