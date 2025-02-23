package graphmap

type Config struct {
	Airstrip         string   `mapstructure:"airstrip"`
	Airport          string   `mapstructure:"airport"`
	AircraftParkings []string `mapstructure:"aircraft_parkings"`
}
