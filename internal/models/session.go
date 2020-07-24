package models

import "time"

type Session struct {
	Nickname string `json:"nickname"`
	Cookie string `json:"cookie"`
	Expiration time.Time `json:"expiration,omitempty"`
}
