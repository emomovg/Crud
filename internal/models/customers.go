package models

import "time"

type Customer struct {
	ID      int64
	Name    string
	Phone   string
	Active  bool
	Created time.Time
}
