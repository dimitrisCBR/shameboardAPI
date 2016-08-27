package model

type Shame struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Source User `json:"source"`
	Destination User `json:"destination"`
}

type Shames []Shame