package model

type District struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Regency Regency `json:"regency"`
}
