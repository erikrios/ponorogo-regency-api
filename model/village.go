package model

type Village struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	District District `json:"district"`
}
