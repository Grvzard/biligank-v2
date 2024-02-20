package config

import "time"

type Config struct {
	Stage1 time.Time
}

var GlobalConfig Config
