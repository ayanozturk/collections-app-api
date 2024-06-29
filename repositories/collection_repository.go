package repositories

import (
	"collections/database"
	"collections/models"
	"log"
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
		log.Printf("Error fetching collections: %v", err)
		return nil, err
	}

	// Fetch items for each collection
	for i, collection := range collections {
		var items []models.Item
		//collections[i].Items = items
		err := database.DB.Select(&items, "SELECT id, collection_id, name, description FROM items WHERE collection_id = ?", collection.ID)
		if err != nil {
			log.Printf("Error fetching items for collection %s: %v", collection.ID, err)
			return nil, err
		}

		if len(items) == 0 {
			collections[i].Items = []models.Item{}
		} else {
			collections[i].Items = items
		}
	}

	if len(collections) == 0 {
		return nil, nil
	}

	return collections, nil
}
