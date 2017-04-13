package datastore

import (
	"fmt"

	boltdb "github.com/boltdb/bolt"
)

// Store is the globally accessible DataStore
var Store *DataStore

// Connect connects to the database
func Connect(path string) error {
	fmt.Println("Initializing data stores")

	db, err := boltdb.Open(path, 0644, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to DB")

	ds := initialDataStore(db)

	Store = ds.NewDataStore()

	return nil
}
