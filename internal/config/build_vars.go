package config

import "os"

var (
	APIURL = os.Getenv("API_URL")
	NATSURL = os.Getenv("NATS_URL")
	FRONTURL = os.Getenv("FRONT_URL")
)