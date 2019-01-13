package models

import "time"

type RequestModelUser struct {
	UUID  string
	Login string
}

type DBModelUser struct {
	UUID  string
	Login string
	Date  time.Time
}
