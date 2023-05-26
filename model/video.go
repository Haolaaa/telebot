package model

import (
	"time"
)

type VideoChanges struct {
	ID        uint32
	Title     string
	PlayUrl   string    `mapstructure:"play_url"`
	DownUrl   string    `mapstructure:"down_url"`
	CreatedAt time.Time `mapstructure:"created_at"`
	UpdatedAt time.Time `mapstructure:"updated_at"`
}
