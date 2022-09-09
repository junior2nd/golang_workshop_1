package model

import "time"

type Product struct {
	ID       uint
	Name     string
	Stock    int64
	Price    float64
	Image    string
	CreateAt time.Time
}
