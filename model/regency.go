package model

type Regency struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Province Province `json:"province"`
}
