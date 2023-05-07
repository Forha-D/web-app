package datastore

import "github.com/Forha-D/web-app/Web-app/entities"

type Animal interface {
	Get(id int64) ([]entities.Animal, error)
	Create (entities.Animal) (entities.Animal, error)
}

