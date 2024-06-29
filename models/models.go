package models

type Collection struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Items []Item `json:"items"`
}

type Item struct {
	ID           string `json:"id"`
	CollectionID string `json:"collection_id"`
	Name         string `json:"name"`
}
