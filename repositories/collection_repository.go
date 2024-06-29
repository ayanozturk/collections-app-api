package repositories

import (
	"collections/database"
	"collections/models"
)

type CollectionRepository interface {
	GetCollections() ([]models.Collection, error)
}

type repository struct{}

func NewCollectionRepository() CollectionRepository {
	return &repository{}
}

func (r *repository) GetCollections() ([]models.Collection, error) {
	var collections []models.Collection
	err := database.DB.Select(&collections, "SELECT id, name FROM collections")
	if err != nil {
		return nil, err
	}

	// Fetch items for each collection
	for i, collection := range collections {
		var items []models.Item
		err := database.DB.Select(&items, "SELECT id, collection_id, name, description FROM items WHERE collection_id = ?", collection.ID)
		if err != nil {
			return nil, err
		}
		collections[i].Items = items
	}

	if len(collections) == 0 {
		return nil, nil
	}

	return collections, nil
}
