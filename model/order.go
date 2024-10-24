package model

import "time"

type Order struct {
	Id    int
	Time  time.Time
	Items []Food
	Total int
}
